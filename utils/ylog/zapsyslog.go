package ylog

//
// import (
//	"log/syslog"
//
//	"github.com/pkg/errors"
//	"go.uber.org/zap/zapcore"
// )
//
// //
// // Core
// //
// // core core
// type core struct {
//	zapcore.LevelEnabler
//
//	encoder zapcore.Encoder
//	writer  *syslog.Writer
//
//	fields []zapcore.Field
// }
//
// //
// //// newCore newCore
// //func newCore(enab zapcore.LevelEnabler, encoder zapcore.Encoder, writer *syslog.Writer) *core {
// //	return &core{
// //		LevelEnabler: enab,
// //		encoder:      encoder,
// //		writer:       writer,
// //	}
// //}
//
// // With With
// func (core *core) With(fields []zapcore.Field) zapcore.Core {
//	// Clone core.
//	clone := *core
//
//	// Clone encoder.
//	clone.encoder = core.encoder.Clone()
//
//	// append fields.
//	for i := range fields {
//		fields[i].AddTo(clone.encoder)
//	}
//	// Done.
//	return &clone
// }
//
// // Check Check
// func (core *core) Check(entry zapcore.Entry, checked *zapcore.CheckedEntry) *zapcore.CheckedEntry {
//	if core.Enabled(entry.Level) {
//		return checked.AddCore(entry, core)
//	}
//	return checked
// }
//
// // Write Write
// func (core *core) Write(entry zapcore.Entry, fields []zapcore.Field) error {
//	// Generate the message.
//	buffer, err := core.encoder.EncodeEntry(entry, fields)
//	if err != nil {
//		return errors.Wrap(err, "failed to encode log entry")
//	}
//
//	message := buffer.String()
//
//	// Write the message.
//	switch entry.Level {
//	case zapcore.DebugLevel:
//		return core.writer.Debug(message)
//
//	case zapcore.InfoLevel:
//		return core.writer.Info(message)
//
//	case zapcore.WarnLevel:
//		return core.writer.Warning(message)
//
//	case zapcore.ErrorLevel:
//		return core.writer.Err(message)
//
//	case zapcore.DPanicLevel:
//		return core.writer.Crit(message)
//
//	case zapcore.PanicLevel:
//		return core.writer.Crit(message)
//
//	case zapcore.FatalLevel:
//		return core.writer.Crit(message)
//
//	default:
//		return errors.Errorf("unknown log level: %v", entry.Level)
//	}
// }
//
// // Sync Sync
// func (core *core) Sync() error {
//	return nil
// }
