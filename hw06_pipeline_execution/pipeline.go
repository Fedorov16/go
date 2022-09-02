package hw06pipelineexecution

type (
	In  = <-chan interface{}
	Out = In
	Bi  = chan interface{}
)

type Stage func(in In) (out Out)

func ExecutePipeline(in In, done In, stages ...Stage) Out {
	for _, fn := range stages {
		in = doMath(fn(in), done)
	}

	return in
}

func doMath(in In, done In) Out {
	out := make(Bi, len(in))
	go runGo(out, in, done)

	return out
}

func runGo(out Bi, in In, done In) {
	defer close(out)
	for {
		select {
		case <-done:
			return
		case res, ok := <-in:
			if !ok {
				return
			}
			out <- res
		}
	}
}
