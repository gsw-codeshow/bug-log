package main

import (
	"fmt"
)

type condValue struct {
	cond   *Condition
	exprs  string
	isCond bool
}

type Condition struct {
	params []condValue
}

func (c Condition) clone() *Condition {
	params := make([]condValue, len(c.params))
	copy(params, c.params)
	c.params = params
	return &c
}

func (c Condition) And(expr string) *Condition {
	c.params = append(c.params, condValue{exprs: expr})
	return &c
}

func (c *Condition) AndCond(cond *Condition) *Condition {
	c = c.clone()
	if c == cond {
		panic(fmt.Errorf("<Condition.AndCond> cannot use self as sub cond"))
	}
	if cond != nil {
		c.params = append(c.params, condValue{cond: cond, isCond: true})
	}
	return c
}

func getCondSQL(cond *Condition) (where string) {
	if cond == nil {
		return
	}
	for _, p := range cond.params {
		if p.isCond {
			w := getCondSQL(p.cond)
			if w != "" {
				w = fmt.Sprintf("( %s ) ", w)
			}
			where += w
		} else {
			exprs := p.exprs
			where += exprs
		}
	}
	return
}

func main() {
	cond := &Condition{}
	cond = cond.AndCond(cond.And("1"))
	cond = cond.AndCond(cond.And("2"))
	cond = cond.AndCond(cond.And("3"))
	cond = cond.AndCond(cond.And("4"))

	cycleFlag := false
	var hasCycle func(*Condition)
	hasCycle = func(c *Condition) {
		if nil == c || cycleFlag {
			return
		}
		condPointMap := make(map[string]bool)
		condPointMap[fmt.Sprintf("%p", c)] = true
		for _, p := range c.params {
			if p.isCond {
				adr := fmt.Sprintf("%p", p.cond)
				if condPointMap[adr] {
					// self as sub cond was cycle
					cycleFlag = true
					return
				} else {
					condPointMap[adr] = true
				}
			}
		}
		for _, p := range c.params {
			if p.isCond {
				// check next
				hasCycle(p.cond)
			}
		}
		return
	}
	hasCycle(cond)
	fmt.Println(cycleFlag)
	return
}
