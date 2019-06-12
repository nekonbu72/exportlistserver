package main

import (
	"errors"
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
	if p.HasQuery() == false {
		return nil, errors.New("URL has no query")
	}

	mailSetting, err := mailg.NewSetting(path.Join(settingDir, mail))
	if err != nil {
		return nil, err
	}

	sinceSlice, ok := p.Query["since"]
	if ok == false {
		return nil, errors.New("No since in query")
	}
	mailSetting.Criteria.Duration.Since = sinceSlice[0]

	beforeSlice, ok := p.Query["before"]
	if ok == false {
		return nil, errors.New("No before in query")
	}
	mailSetting.Criteria.Duration.Before = beforeSlice[0]

	excelSetting, err := exportlistmapping.NewSetting(path.Join(settingDir, excel))
	if err != nil {
		return nil, err
	}

	mailClient, err := mailg.Login(mailSetting.ConnInfo)
	if err != nil {
		return nil, err
	}
	defer mailClient.Logout()

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
