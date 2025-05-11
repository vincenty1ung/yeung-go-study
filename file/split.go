package file

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"

	"golang.org/x/sync/semaphore"
)

const (
	ChunkSize     = 10 * 1024 * 1024 // 1MB
	MaxGoroutines = 4                // 可根据 CPU 数量调整
)

type ChunkResult struct {
	Filename string
	Hash     string
	Err      error
}

func SplitFileConcurrent(ctx context.Context, filePath, outputDir string, chunkSize int) error {
	inputFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	var wg sync.WaitGroup
	sem := semaphore.NewWeighted(int64(MaxGoroutines))
	resultChan := make(chan ChunkResult)
	if chunkSize < 0 {
		chunkSize = ChunkSize
	}
	part := 0
	buffer := make([]byte, chunkSize)

	for {
		n, err := inputFile.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		// 拷贝当前块数据（避免后续被覆盖）
		chunkData := make([]byte, n)
		copy(chunkData, buffer[:n])
		chunkName := fmt.Sprintf("part_%d", part)
		part++
		wg.Add(1)

		go func(data []byte, name string) {
			defer wg.Done()
			sem.Acquire(ctx, 1)
			defer sem.Release(1)

			filePath := filepath.Join(outputDir, name)
			err := os.WriteFile(filePath, data, 0644)
			if err != nil {
				resultChan <- ChunkResult{name, "", err}
				return
			}

			hash := sha256.Sum256(data)
			resultChan <- ChunkResult{name, hex.EncodeToString(hash[:]), nil}
		}(chunkData, chunkName)
	}

	// 收集所有结果
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 写入 hash 文件
	hashFile, err := os.Create(filepath.Join(outputDir, "hashes.sha256"))
	if err != nil {
		return err
	}
	defer hashFile.Close()

	for result := range resultChan {
		if result.Err != nil {
			return fmt.Errorf("分片 %s 出错: %v", result.Filename, result.Err)
		}
		_, err := fmt.Fprintf(hashFile, "%s  %s\n", result.Hash, result.Filename)
		if err != nil {
			return err
		}
	}

	fmt.Println("分片完成，所有 SHA256 哈希写入成功。")
	return nil
}

func readHashFile(path string) (map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hashes := make(map[string]string)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "  ")
		if len(parts) == 2 {
			hashes[parts[1]] = parts[0]
		}
	}
	return hashes, scanner.Err()
}

func MergeChunksConcurrent(ctx context.Context, chunkDir, outputFile string) error {
	hashes, err := readHashFile(filepath.Join(chunkDir, "hashes.sha256"))
	if err != nil {
		return err
	}

	files, err := os.ReadDir(chunkDir)
	if err != nil {
		return err
	}

	var chunkFiles []string
	for _, file := range files {
		if strings.HasPrefix(file.Name(), "part_") {
			chunkFiles = append(chunkFiles, file.Name())
		}
	}
	// sort.Strings(chunkFiles)

	sort.Slice(
		chunkFiles, func(i, j int) bool {
			getNum := func(s string) int {
				var num int
				fmt.Sscanf(filepath.Base(s), "part_%d", &num)
				return num
			}
			return getNum(chunkFiles[i]) < getNum(chunkFiles[j])
		},
	)

	type Chunk struct {
		Order int
		Data  []byte
		Err   error
	}

	resultChan := make(chan Chunk, len(chunkFiles))
	sem := semaphore.NewWeighted(int64(MaxGoroutines))
	var wg sync.WaitGroup

	for idx, filename := range chunkFiles {
		wg.Add(1)
		go func(order int, fname string) {
			defer wg.Done()
			sem.Acquire(ctx, 1)
			defer sem.Release(1)

			data, err := os.ReadFile(filepath.Join(chunkDir, fname))
			if err != nil {
				resultChan <- Chunk{order, nil, err}
				return
			}

			expected := hashes[fname]
			actual := sha256.Sum256(data)
			if expected != hex.EncodeToString(actual[:]) {
				resultChan <- Chunk{order, nil, fmt.Errorf("哈希不匹配：%s", fname)}
				return
			}

			resultChan <- Chunk{order, data, nil}
		}(idx, filename)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// 收集并排序
	chunkData := make([][]byte, len(chunkFiles))
	for chunk := range resultChan {
		if chunk.Err != nil {
			return chunk.Err
		}
		chunkData[chunk.Order] = chunk.Data
	}

	outFile, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer outFile.Close()

	for _, data := range chunkData {
		_, err := outFile.Write(data)
		if err != nil {
			return err
		}
	}

	fmt.Printf("文件成功合并并校验：%s\n", outputFile)

	// ✅ 合并成功后删除分片和哈希文件
	for _, fname := range chunkFiles {
		err := os.Remove(filepath.Join(chunkDir, fname))
		if err != nil {
			fmt.Printf("无法删除分片 %s：%v\n", fname, err)
		}
	}
	err = os.Remove(filepath.Join(chunkDir, "hashes.sha256"))
	if err != nil {
		fmt.Printf("无法删除哈希文件：%v\n", err)
	}

	fmt.Println("原始分片和哈希文件已删除。")
	return nil
}

func CalcFileSHA256(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		return "ERROR"
	}
	defer file.Close()

	hasher := sha256.New()
	io.Copy(hasher, file)
	return hex.EncodeToString(hasher.Sum(nil))
}
