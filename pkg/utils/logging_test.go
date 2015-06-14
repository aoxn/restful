package utils

import (
    "testing"

)


/*
测试格式：func TestXxx (t *testing.T),Xxx部分可以为任意的字母数字的组合，但是首字母不能是小写字母[a-z]，例如Testintdiv是错误的函数名。
函数中通过调用testing.T的Error, Errorf, FailNow, Fatal, FatalIf方法，说明测试不通过，调用Log方法用来记录测试的信息。

 */

func TestOnValueChange1(t *testing.T) {
    objModule := logmodule{
        bIsFirst:true,
        strFile:"",
        strSize:"",
        vecModule:[]string{LOG_MODULE_APISVR,
            LOG_MODULE_SCHED, LOG_MODULE_ARCH, LOG_MODULE_MASTERCORE, LOG_MODULE_SLAVECORE, LOG_MODULE_SLAVERT},
    }
    bIsOk:=objModule.OnValueChange("abc", "old1", "new1")
    if !bIsOk{
        t.Error("TestOnValueChange1 not correct.")
        t.FailNow()
    }
}


func TestOnValueChange2(t *testing.T) {
    objModule := logmodule{
        bIsFirst:true,
        strFile:"",
        strSize:"",
        vecModule:[]string{LOG_MODULE_APISVR,
            LOG_MODULE_SCHED, LOG_MODULE_ARCH, LOG_MODULE_MASTERCORE, LOG_MODULE_SLAVECORE, LOG_MODULE_SLAVERT},
    }

    bIsOk:=objModule.OnValueChange(LOGGING_PROPERTY_FILE, "old2", "/data/yy/log/rds_container_api_d/rds_container_api_d.log")
    if !bIsOk{
        t.Error("objModule.OnValueChange(LOGGING_PROPERTY_FILE not correct.")
        t.FailNow()
    }

    bIsOk=objModule.OnValueChange(LOGGING_PROPERTY_ROTATE_SIZE, "old2", 1024)
    if bIsOk{
        t.Error("objModule.OnValueChange(LOGGING_PROPERTY_ROTATE_SIZE not correct.")
        t.FailNow()
    }


}



//wuji 测试函数一般是以“Test”为名称前缀并有一个类型为“testing.T”的参数声明的函数.
func TestOnValueChange3(t *testing.T) {

    objModule := logmodule{
        bIsFirst:true,
        strFile:"",
        strSize:"",
        vecModule:[]string{LOG_MODULE_APISVR,
            LOG_MODULE_SCHED,LOG_MODULE_ARCH,LOG_MODULE_MASTERCORE,LOG_MODULE_SLAVECORE,LOG_MODULE_SLAVERT},
    }

    bIsOk:=objModule.OnValueChange(LOGGING_PROPERTY_FILE, "old2", "/data/yy/log/rds_container_api_d/rds_container_api_d.log")
    if !bIsOk{
        t.Error("objModule.OnValueChange(LOGGING_PROPERTY_FILE not correct.")
        t.FailNow()
    }

    bIsOk=objModule.OnValueChange(LOGGING_PROPERTY_ROTATE_SIZE, "old2", "1024")
    if !bIsOk{
        t.Error("objModule.OnValueChange(LOGGING_PROPERTY_ROTATE_SIZE not correct.")
        t.FailNow()
    }

    bIsOk=objModule.OnValueChange("rds.logging.level", "old2", "info")
    if bIsOk{
        t.Error("rds.logging.level not correct.")
        t.FailNow()
    }






}


func TestOnValueChange4(t *testing.T) {

    objModule := logmodule{
        bIsFirst:true,
        strFile:"",
        strSize:"",
        vecModule:[]string{LOG_MODULE_APISVR,
            LOG_MODULE_SCHED,LOG_MODULE_ARCH,LOG_MODULE_MASTERCORE,LOG_MODULE_SLAVECORE,LOG_MODULE_SLAVERT},
    }


    bIsOk:=objModule.OnValueChange(LOGGING_PROPERTY_FILE, "old2", "/data/yy/log/rds_container_api_d/rds_container_api_d.log")
    if !bIsOk{
        t.Error("objModule.OnValueChange(LOGGING_PROPERTY_FILE not correct.")
        t.FailNow()
    }

    bIsOk=objModule.OnValueChange(LOGGING_PROPERTY_ROTATE_SIZE, "old2", "1024")
    if !bIsOk{
        t.Error("objModule.OnValueChange(LOGGING_PROPERTY_ROTATE_SIZE not correct.")
        t.FailNow()
    }

    bIsOk=objModule.OnValueChange("apisvr.logging.level", "old2", "info")
    if !bIsOk{
        t.Error("apisvr.logging.level not correct.")
        t.FailNow()
    }


    objLogger:=GetLogger("apisvr")
    objLogger.logger.Info("api svr is right")

}



