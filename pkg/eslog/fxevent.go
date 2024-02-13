package eslog

import (
	"log/slog"
	"strings"

	"go.uber.org/fx/fxevent"
)

// LogEvent implements fxevent.Logger
func (l *Logger) LogEvent(event fxevent.Event) {
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.Trace("OnStart hook executing",
			slog.String("callee", e.FunctionName),
			slog.String("caller", e.CallerName),
		)
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.Error("OnStart hook failed",
				slog.String("callee", e.FunctionName),
				slog.String("caller", e.CallerName),
				slog.String("err", e.Err.Error()),
			)
		} else {
			l.Trace("OnStart hook executed",
				slog.String("callee", e.FunctionName),
				slog.String("caller", e.CallerName),
				slog.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.OnStopExecuting:
		l.Trace("OnStop hook executing",
			slog.String("callee", e.FunctionName),
			slog.String("caller", e.CallerName),
		)
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.Error("OnStop hook failed",
				slog.String("callee", e.FunctionName),
				slog.String("caller", e.CallerName),
				slog.String("err", e.Err.Error()),
			)
		} else {
			l.Trace("OnStop hook executed",
				slog.String("callee", e.FunctionName),
				slog.String("caller", e.CallerName),
				slog.String("runtime", e.Runtime.String()),
			)
		}
	case *fxevent.Supplied:
		if e.Err != nil {
			l.Error("error encountered while applying options",
				slog.String("type", e.TypeName),
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("module", e.ModuleName), //moduleField(e.ModuleName),
				slog.String("err", e.Err.Error()),
			)
		} else {
			l.Trace("supplied",
				slog.String("type", e.TypeName),
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("module", e.ModuleName), //moduleField(e.ModuleName),
			)
		}
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.Trace("provided",
				slog.String("constructor", e.ConstructorName),
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("module", e.ModuleName), //moduleField(e.ModuleName),
				slog.String("type", rtype),
				slog.Bool("private", e.Private), //maybeBool("private", e.Private),
			)
		}
		if e.Err != nil {
			l.Error("error encountered while applying options",
				slog.String("module", e.ModuleName), //moduleField(e.ModuleName),
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("err", e.Err.Error()),
			)
		}
	case *fxevent.Replaced:
		for _, rtype := range e.OutputTypeNames {
			l.Trace("replaced",
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("module", e.ModuleName), //moduleField(e.ModuleName),
				slog.String("type", rtype),
			)
		}
		if e.Err != nil {
			l.Error("error encountered while replacing",
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("module", e.ModuleName), //moduleField(e.ModuleName),
				slog.String("err", e.Err.Error()),
			)
		}
	case *fxevent.Decorated:
		for _, rtype := range e.OutputTypeNames {
			l.Trace("decorated",
				slog.String("decorator", e.DecoratorName),
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("module", e.ModuleName), //moduleField(e.ModuleName),
				slog.String("type", rtype),
			)
		}
		if e.Err != nil {
			l.Error("error encountered while applying options",
				slog.Any("stacktrace", e.StackTrace),
				slog.Any("moduletrace", e.ModuleTrace),
				slog.String("module", e.ModuleName), //moduleField(e.ModuleName),
				slog.String("err", e.Err.Error()),
			)
		}
	case *fxevent.Run:
		if e.Err != nil {
			l.Error("error returned",
				slog.String("name", e.Name),
				slog.String("kind", e.Kind),
				slog.String("module", e.ModuleName), //moduleField(e.ModuleName),
				slog.String("err", e.Err.Error()),
			)
		} else {
			l.Trace("run",
				slog.String("name", e.Name),
				slog.String("kind", e.Kind),
				slog.String("module", e.ModuleName), //moduleField(e.ModuleName),
			)
		}
	case *fxevent.Invoking:
		// Do not log stack as it will make logs hard to read.
		l.Trace("invoking",
			slog.String("function", e.FunctionName),
			slog.String("module", e.ModuleName), //moduleField(e.ModuleName),
		)
	case *fxevent.Invoked:
		if e.Err != nil {
			l.Error("invoke failed",
				slog.String("err", e.Err.Error()),
				slog.Any("stack", stackFormat(e.Trace)),
				slog.String("function", e.FunctionName),
				slog.String("module", e.ModuleName), //moduleField(e.ModuleName),
			)
		}
	case *fxevent.Stopping:
		l.Info("received signal",
			slog.String("signal", strings.ToUpper(e.Signal.String())),
		)
	case *fxevent.Stopped:
		if e.Err != nil {
			l.Error("stop failed",
				slog.String("err", e.Err.Error()),
			)
		}
	case *fxevent.RollingBack:
		l.Error("start failed, rolling back",
			slog.String("err", e.StartErr.Error()),
		)
	case *fxevent.RolledBack:
		if e.Err != nil {
			l.Error("rollback failed",
				slog.String("err", e.Err.Error()),
			)
		}
	case *fxevent.Started:
		if e.Err != nil {
			l.Error("start failed",
				slog.String("err", e.Err.Error()),
			)
		} else {
			l.Info("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.Error("custom logger initialization failed",
				slog.String("err", e.Err.Error()),
			)
		} else {
			l.Trace("initialized custom fxevent.Logger",
				slog.String("function", e.ConstructorName),
			)
		}
	}
}
