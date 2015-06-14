/*
  this module provide a machanism to access the 
  property file for the application
*/
package utils

import (
   "fmt"
   "os"
   "time"
   "reflect"
   "github.com/pelletier/go-toml"
)

const (
   CONFIG_SCAN_INTERVAL = 1 * time.Second
)

type ConfWatcher interface {
    OnValueChange(key string, oldval interface{}, newval interface{}) bool
}

type ConfMgr interface {
   //register a watcher for a specific key.
   //key changes will be notified through ConfWatcher interface
   RegisterWatcher(k string, w ConfWatcher)
   Refresh()
}

type PropertyFileMgr struct {
   fname string
   oldvals map[string]interface{}
   newvals map[string]interface{}
   watchers map[string] []ConfWatcher
}

func NewPropertyFileMgr(fname string) *PropertyFileMgr {
   m := &PropertyFileMgr{}
   m.fname = fname
   m.oldvals = make(map[string]interface{})
   m.newvals = make(map[string]interface{})
   m.watchers = make(map[string] []ConfWatcher)
   return m
}

func (m *PropertyFileMgr) RegisterWatcher(key string, w ConfWatcher) {
   if val, ok := m.watchers[key]; ok {//val是一个ConfWatcher数组
      val = append(val, w)
      m.watchers[key] = val
   }else{
      val = append(val, w)
      m.watchers[key] = val
   }
}

//update m.newvals
//use m.newvals to decouple the readProperty with scan loop for unit test
func (m *PropertyFileMgr) readProperty() {
   fmt.Println("open config file [", m.fname, "] start")

   tree, err := toml.LoadFile(m.fname)
   if err!=nil {
      fmt.Fprintln(os.Stderr, "Fail to open config file [", m.fname, "]")
      return
   }
   vals := make(map[string]interface{})
   for k, _:=range m.watchers {
       fmt.Println("k:",k)
      if tree.Has(k) {
         vals[k] = tree.Get(k)
          fmt.Println("k:",k,"vals[k]:",vals[k])
      }
   }
   m.newvals = vals
    fmt.Println("open config file [", m.fname, "] ok")
}

//true for equal, false for not equal
func compareEqual(v1 interface{}, v2 interface{}) bool {
   return reflect.DeepEqual(v1, v2)
}

//scan the newvals and notify the changes to 
//observers
func (m *PropertyFileMgr) compareAndNotify() {
    fmt.Println("compareAndNotify start")

   for k,v := range m.newvals { //compare with oldvals
       fmt.Println("compareAndNotify ",k, v)
      if !compareEqual(v, m.oldvals[k]) {
          fmt.Println("compareAndNotify !compareEqual",k, v)
         //not equal
         if watchers, ok := m.watchers[k]; ok {
            for _, w := range watchers {
                fmt.Println("compareAndNotify ",m.oldvals[k], m.newvals[k])
               w.OnValueChange(k, m.oldvals[k], m.newvals[k])
            }
         }
      }
   }
   m.oldvals = m.newvals
    fmt.Println("compareAndNotify ok")
}

func (m *PropertyFileMgr) MainLoop() {
   for {
      time.Sleep(CONFIG_SCAN_INTERVAL)
      m.readProperty()
      m.compareAndNotify()
   }
}

func (m *PropertyFileMgr) Refresh() {
        m.readProperty()
        m.compareAndNotify()
}
