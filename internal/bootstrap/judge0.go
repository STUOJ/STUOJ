package bootstrap

import (
	"STUOJ/external/judge0"
	"STUOJ/internal/service/language"
	"STUOJ/utils"
	"log"
)

func initJudge0() {
	var err error
	err = judge0.InitJudge()
	if err != nil {
		log.Println(err)
		log.Println("初始化评测机失败！")
		return
	}

	/*	err = InitJudgePrintInfo()
		if err != nil {
			log.Println(err)
			log.Println("Init judge0 failed!")
			return
		}
	*/

	err = InitJudgeLanguages()
	if err != nil {
		log.Println(err)
		log.Println("初始化评测机语言失败！")
		return
	}

}

// 初始化评测机语言
func InitJudgeLanguages() error {
	// 读取评测机语言列表
	languages, err := judge0.GetLanguage()
	if err != nil {
		return err
	}

	/*	log.Println("Judge support languages:")
		for k, v := range languages {
			log.Println(k, v)
		}
	*/

	// 清空数据库语言表
	err = language.DeleteAll()
	if err != nil {
		log.Println(err)
	}

	// 插入数据库语言表
	for _, l := range languages {
		//log.Println(v)

		// 初始化对象
		// language := entity.Language{
		// 	Id:   uint64(v["id"].(float64)),
		// 	Name: v["name"].(string),
		// }

		_, err := language.Insert(l)
		if err != nil {
			return err
		}
	}

	return nil
}

// 打印评测机信息
func InitJudgePrintInfo() error {
	config, err := judge0.GetConfigInfo()
	if err != nil {
		return err
	}
	if configtmp, err := utils.PrettyStruct(config); err != nil {
		log.Println("Struct formatting failed:", err)
		log.Println("Judge config info:", config)
	} else {
		log.Println("Judge config info:", configtmp)
	}

	system, err := judge0.GetSystemInfo()
	if err != nil {
		return err
	}
	if systemtmp, err := utils.PrettyStruct(system); err != nil {
		log.Println("Struct formatting failed:", err)
		log.Println("Judge system info:", system)
	} else {
		log.Println("Judge system info:", systemtmp)
	}

	/*	statistics, err := judge0.GetStatistics()
		if err != nil {
			return err
		}
		if statstmp, err := utils.PrettyStruct(statistics); err != nil {
			log.Println("Struct formatting failed:", err)
			log.Println("Judge statistics:", statistics)
		} else {
			log.Println("Judge statistics:", statstmp)
		}
	*/

	/*	about, err := judge0.GetAbout()
		if err != nil {
			return err
		}
		log.Println("Judge about:", about)
	*/

	workers, err := judge0.GetWorkers()
	if err != nil {
		return err
	}
	log.Println("Judge workers:")
	for _, worker := range workers {
		if workerstmp, err := utils.PrettyStruct(worker); err != nil {
			log.Println("Struct formatting failed:", err)
			log.Println(worker)
		} else {
			log.Println(workerstmp)
		}
	}

	/*	license, err := judge0.GetLicense()
		if err != nil {
			return err
		}
		log.Println("Judge license:", license)
	*/

	/*	isolate, err := judge0.GetIsolate()
		if err != nil {
			return err
		}
		log.Println("Judge isolate:", isolate)
	*/

	version, err := judge0.GetVersion()
	if err != nil {
		return err
	}
	log.Println("Judge version:", version)

	return nil
}
