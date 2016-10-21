package lang

import (
	"github.com/ONSdigital/dp-frontend-renderer/assets"
	"github.com/nicksnyder/go-i18n/i18n"
)

func Load(lang ...string) error {
	for _, l := range lang {
		err := load(l)
		if err != nil {
			return err
		}
	}

	return nil
}

func load(lang string) error {
	b, err := assets.Asset("lang/" + lang + ".json")
	if err != nil {
		return err
	}

	err = i18n.ParseTranslationFileBytes("lang/"+lang+".json", b)
	if err != nil {
		return err
	}

	return nil
}

func Get(lang string) (i18n.TranslateFunc, error) {
	return i18n.Tfunc(lang)
}
