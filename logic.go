package main

import (
	"path"

	"github.com/nekonbu72/xemlsx"

	"github.com/nekonbu72/exportlistmapping"
	"github.com/nekonbu72/mailg"
)

const (
	settingDir = "settings"
	excel      = "excel.json"
	mail       = "mail.json"
)

func businessLogic(p *Path) ([]*exportlistmapping.Data, error) {
	mailSetting, err := mailg.NewSetting(path.Join(settingDir, mail))
	if err != nil {
		return nil, err
	}

	excelSetting, err := exportlistmapping.NewSetting(path.Join(settingDir, excel))
	if err != nil {
		return nil, err
	}

	mailClient, err := mailg.Login(mailSetting.ConnInfo)
	if err != nil {
		return nil, err
	}
	defer mailClient.Logout()

	if p.HasQuery() == false {
		//
		mailSetting.Criteria.Duration.Since = "2019-06-11 JST"
		mailSetting.Criteria.Duration.Before = "2019-06-12 JST"
		//
	}

	done := make(chan interface{})
	defer close(done)

	attacmentStream := mailClient.FetchAttachment(done, mailSetting.Criteria)
	xlsxStream := xemlsx.ToXLSX(done, attacmentStream)
	data, err := exportlistmapping.Fetch(excelSetting, xlsxStream)
	if err != nil {
		return nil, err
	}
	return data, nil
}
