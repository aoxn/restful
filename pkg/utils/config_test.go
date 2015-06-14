package utils

import (
   "testing"
   "os"
   "io/ioutil"
)


//wuji 测试函数一般是以“Test”为名称前缀并有一个类型为“testing.T”的参数声明的函数.
func TestCompareValue(t *testing.T) {
   a := "123"
   b := 123
   if compareEqual(a, b) {
     t.Error("compare Value not correct.")
     t.FailNow()
   }
}

func TestCompareValue1(t *testing.T) {
   a := 123
   b := 123
   if !compareEqual(a, b) {
     t.Error("compare Value not correct.")
     t.FailNow()
   }
}

func TestCompareValue2(t *testing.T) {
   a := 123
   if compareEqual(a, nil) {
     t.Error("compare Value not correct.")
     t.FailNow()
   }
}

type MockWatcher struct {
   key string
   oldval interface{}
   newval interface{}
}

func (w *MockWatcher) OnValueChange(k string, v1 interface{}, v2 interface{}) bool {
   w.key = k
   w.oldval = v1
   w.newval = v2
   return true
}

func TestRegisterWatcher(t *testing.T) {
   m := NewPropertyFileMgr("")
   w1 := &MockWatcher{}
   w2 := &MockWatcher{}
   m.RegisterWatcher("testkey", w1)
   m.RegisterWatcher("testkey", w2)

   var w1ok, w2ok bool
   for _, w := range m.watchers["testkey"] {
      if w == w1 {
         w1ok = true
         continue
      }
      if w == w2 {
         w2ok = true
         continue
      }
      t.Error("more than 2 watchers were registered!")
      t.FailNow()
   }
   if !(w1ok && w2ok) {
      t.Error("RegisterWatcher not correct.")
      t.FailNow()
   }
}

func TestCompareAndNotify(t *testing.T) {
   m := NewPropertyFileMgr("")
   w1 := &MockWatcher{}
   m.RegisterWatcher("testkey", w1)
   m.oldvals["testkey"] = 123
   m.newvals["testkey"] = 123
   m.compareAndNotify()

   if w1.key != "" {
      t.Error("compareAndNotify not correct. w1.key=", w1.key)
      t.FailNow()
   }
}

func TestCompareAndNotify1(t *testing.T) {
   m := NewPropertyFileMgr("")
   w1 := &MockWatcher{}
   m.RegisterWatcher("testkey", w1)
   //m.oldvals["testkey"] = 123
   m.newvals["testkey"] = 123
   m.compareAndNotify()

   if w1.oldval != nil || w1.newval.(int) != 123  {
      t.Error("compareAndNotify not correct")
      t.FailNow()
   }
}

func delFile(f string) {
   os.RemoveAll(f)
}

func createFile(f, c string) {
   ioutil.WriteFile(f, []byte(c), 0777)
}

func TestReadProperty(t *testing.T) {
   delFile("test.ini")
   var c =
` 
[sigma]
ip = "123"
`
   createFile("test.ini", c)
   m := NewPropertyFileMgr("test.ini")
   m.watchers["sigma.ip"] = nil
   m.readProperty()
   v, ok:=m.newvals["sigma.ip"];
   if ok {
      if v.(string)=="123" {
         return
      }
   }
   t.Error("readProperty not correct.", ok, v)
   t.FailNow()
}

