package zaplog

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	defaultLogMaxSize = 300 // MB
)

// FileLogConfig serializes file log related config in yaml/toml/json.
type FileLogConfig struct {
	// Log filename, leave empty to disable file log.
	Filename string `yaml:"filename" toml:"filename" json:"filename"`
	// Max size for a single file, in MB.
	MaxSize int `yaml:"max-size" toml:"max-size" json:"max-size"`
	// Max log keep days, default is never deleting.
	MaxDays int `yaml:"max-days" toml:"max-days" json:"max-days"`
	// Maximum number of old log files to retain.
	MaxBackups int `yaml:"max-backups" toml:"max-backups" json:"max-backups"`
}

// Config serializes log related config in yaml/toml/json.
type Config struct {
	// Log level.
	Level string `yaml:"level" toml:"level" json:"level"`
	// Log format. one of json, text, or console.
	Format string `yaml:"format" toml:"format" json:"format"`
	// Disable automatic timestamps in output.
	DisableTimestamp bool `yaml:"disable-timestamp" toml:"disable-timestamp" json:"disable-timestamp"`
	// File log config.
	File FileLogConfig `yaml:"file" toml:"file" json:"file"`
	// Development puts the logger in development mode, which changes the
	// behavior of DPanicLevel and takes stacktraces more liberally.
	Development bool `yaml:"development" toml:"development" json:"development"`
	// DisableCaller stops annotating logs with the calling function's file
	// name and line number. By default, all logs are annotated.
	DisableCaller bool `yaml:"disable-caller" toml:"disable-caller" json:"disable-caller"`
	// DisableStacktrace completely disables automatic stacktrace capturing. By
	// default, stacktraces are captured for WarnLevel and above logs in
	// development and ErrorLevel and above in production.
	DisableStacktrace bool `yaml:"disable-stacktrace" toml:"disable-stacktrace" json:"disable-stacktrace"`
	// DisableErrorVerbose stops annotating logs with the full verbose error
	// message.
	DisableErrorVerbose bool `yaml:"disable-error-verbose" toml:"disable-error-verbose" json:"disable-error-verbose"`
	// SamplingConfig sets a sampling strategy for the logger. Sampling caps the
	// global CPU and I/O load that logging puts on your process while attempting
	// to preserve a representative subset of your logs.
	//
	// Values configured here are per-second. See zapcore.NewSampler for details.
	Sampling *zap.SamplingConfig `yaml:"sampling" toml:"sampling" json:"sampling"`
}

// ZapProperties records some information about zap.
type ZapProperties struct {
	Core   zapcore.Core
	Syncer zapcore.WriteSyncer
	Level  zap.AtomicLevel
}

func newZapTextEncoder(cfg *Config) zapcore.Encoder {
	return NewTextEncoder(cfg)
}

func (cfg *Config) buildOptions(errSink zapcore.WriteSyncer) []zap.Option {
	opts := []zap.Option{zap.ErrorOutput(errSink)}

	if cfg.Development {
		opts = append(opts, zap.Development())
	}

	if !cfg.DisableCaller {
		opts = append(opts, zap.AddCaller())
	}

	stackLevel := zap.ErrorLevel
	if cfg.Development {
		stackLevel = zap.WarnLevel
	}
	if !cfg.DisableStacktrace {
		opts = append(opts, zap.AddStacktrace(stackLevel))
	}

	if cfg.Sampling != nil {
		opts = append(opts, zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return zapcore.NewSampler(core, time.Second, int(cfg.Sampling.Initial), int(cfg.Sampling.Thereafter))
		}))
	}
	return opts
}
