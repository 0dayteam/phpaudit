package audit

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/node/expr"
	"github.com/z7zmey/php-parser/php7"
	"log"
	"phpaudit/phpread"
)

type Auditor struct {
	file       []string
	fileManage FileManage
}

func (s *Auditor) scan() {
	ch := make(chan string, 5)
	for i := 0; i < 5; i++ {
		go func(c chan string) {
			for v := range c {
				f, err := phpread.NewPhpFile(v)
				// todo change log method
				if err != nil {
					log.Print(err)
					continue
				}
				errs := f.Parser()
				for _, e := range errs {
					log.Printf("%s", e)
					continue
				}
				//  s.fileManage[v] = f.Code
			}
		}(ch)
	}
	for _, v := range s.file {
		ch <- v
	}
	for len(ch) == 0 {
		close(ch)
	}
}

func (s *Auditor) run(code *php7.Parser) {
	// r := code.GetRootNode().(*node.Root)

}

func (s *Auditor) Explain(n node.Node) {
	switch n.(type) {
	case *expr.Include:
	case *expr.Exit:
	case *expr.Isset:

	}
}
