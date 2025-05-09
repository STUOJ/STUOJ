package bootstrap

import (
	"STUOJ/internal/domain/language"
	"STUOJ/internal/domain/runner"
	judge1 "STUOJ/internal/infrastructure/judge0"
	"STUOJ/internal/infrastructure/persistence/repository/dao"
	"STUOJ/internal/infrastructure/persistence/repository/option"
	"STUOJ/pkg/config"
	"STUOJ/pkg/utils"
	"log"
)

func InitJudge0() {
	var err error
	err = judge1.InitJudge(config.Conf.Judge.Host, config.Conf.Judge.Port, config.Conf.Judge.Token)
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
	languages, err := runner.Runner.GetLanguage()
	if err != nil {
		return err
	}

	oldLang_Map, err := dao.LanguageStore.Select(option.NewQueryOptions())
	if err != nil {
		return err
	}
	oldLangs := make([]language.Language, len(oldLang_Map))
	for i := range oldLang_Map {
		oldLangs[i] = language.Dto(oldLang_Map[i])
	}
	oldLangMap := make(map[string]*language.Language, len(oldLangs))
	for i := range oldLangs {
		oldLangMap[oldLangs[i].Name.String()] = &oldLangs[i]
	}
	for i := range languages {
		if lang, exists := oldLangMap[languages[i].Name]; exists && languages[i].Id != int64(lang.MapId.Value()) {
			lang.MapId.Set(uint32(languages[i].Id))
			if err := lang.Update(); err != nil {
				return err
			}
		} else if !exists {
			log.Printf("找不到：%+v\n", languages[i])
			lang := language.NewLanguage(language.WithName(languages[i].Name), language.WithMapId(uint32(languages[i].Id)))
			if _, err := lang.Create(); err != nil {
				return err
			}
		}
	}
	return nil
}

// 打印评测机信息
func InitJudgePrintInfo() error {
	config, err := judge1.GetConfigInfo()
	if err != nil {
		return err
	}
	if configtmp, err := utils.PrettyStruct(config); err != nil {
		log.Println("Struct formatting failed:", err)
		log.Println("Judge config info:", config)
	} else {
		log.Println("Judge config info:", configtmp)
	}

	system, err := judge1.GetSystemInfo()
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

	workers, err := judge1.GetWorkers()
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

	version, err := judge1.GetVersion()
	if err != nil {
		return err
	}
	log.Println("Judge version:", version)

	return nil
}
