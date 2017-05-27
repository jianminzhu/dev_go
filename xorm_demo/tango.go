
package main

import (
    //"errors"
    "github.com/lunny/tango"
)
 
type EventAction struct {
	tango.Ctx
}

func (c *EventAction) Get() {
	c.Write([]byte("get"))
}

func (c *EventAction) Before() {
	c.Write([]byte("before "))
}

func (c *EventAction) After() {
	c.Write([]byte(" after"))
}

func main() {
   t := tango.New(tango.Static()) 
    t.Use(events.Events())
    t.Get("/", new(EventAction))
    t.Run()
}


//type Action struct {
//    tango.JSON
//}
//
//func (Action) Get() interface{} {
//    if true {
//        return map[string]string{
//            "say": "Hello tango!",
//        }
//    }
//    return errors.New("something error")
//}
//
//func main() {
//    t := tango.Classic()
//    t.Get("/", new(Action))
//    t.Run()
//}