package file

import (
	"context"
	"fmt"
	"testing"
)

func TestSplitFileConcurrentAndMergeChunksConcurrent(t *testing.T) {
	ctx := context.Background()

	err := SplitFileConcurrent(ctx, "20250506.zip", "./chunks", 10*1024*1024)
	if err != nil {
		fmt.Println("分片失败：", err)
		return
	}

	err = MergeChunksConcurrent(ctx, "./chunks", "merged_20250506.zip")
	if err != nil {
		fmt.Println("合并失败：", err)
	}
}

func TestSplitFileConcurrent(t *testing.T) {
	ctx := context.Background()

	err := SplitFileConcurrent(ctx, "20250506.zip", "./chunks", 30*1024*1024)
	if err != nil {
		fmt.Println("分片失败：", err)
		return
	}

}

func TestMergeChunksConcurrent(t *testing.T) {
	ctx := context.Background()

	err := MergeChunksConcurrent(ctx, "./chunks", "merged_20250506.zip")
	if err != nil {
		fmt.Println("合并失败：", err)
	}
}

func TestName1(t *testing.T) {
	originalHash := CalcFileSHA256("20250506.zip")
	mergedHash := CalcFileSHA256("merged_20250506.zip")
	if originalHash != mergedHash {
		fmt.Println("文件合并后与原始文件不一致")
	}
}
