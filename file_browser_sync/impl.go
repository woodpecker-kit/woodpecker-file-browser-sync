package file_browser_sync

import (
	"fmt"
	"github.com/sinlov-go/go-common-lib/pkg/string_tools"
	"github.com/sinlov-go/go-common-lib/pkg/struct_kit"
	"github.com/woodpecker-kit/woodpecker-tools/wd_info"
	"github.com/woodpecker-kit/woodpecker-tools/wd_log"
	"github.com/woodpecker-kit/woodpecker-tools/wd_short_info"
	"path/filepath"
)

func (p *FileBrowserSyncPlugin) ShortInfo() wd_short_info.WoodpeckerInfoShort {
	if p.wdShortInfo == nil {
		info2Short := wd_short_info.ParseWoodpeckerInfo2Short(*p.woodpeckerInfo)
		p.wdShortInfo = &info2Short
	}
	return *p.wdShortInfo
}

// SetWoodpeckerInfo
// also change ShortInfo() return
func (p *FileBrowserSyncPlugin) SetWoodpeckerInfo(info wd_info.WoodpeckerInfo) {
	var newInfo wd_info.WoodpeckerInfo
	_ = struct_kit.DeepCopyByGob(&info, &newInfo)
	p.woodpeckerInfo = &newInfo
	info2Short := wd_short_info.ParseWoodpeckerInfo2Short(newInfo)
	p.wdShortInfo = &info2Short
}

func (p *FileBrowserSyncPlugin) GetWoodPeckerInfo() wd_info.WoodpeckerInfo {
	return *p.woodpeckerInfo
}

func (p *FileBrowserSyncPlugin) OnlyArgsCheck() {
	p.onlyArgsCheck = true
}

func (p *FileBrowserSyncPlugin) Exec() error {
	errLoadStepsTransfer := p.loadStepsTransfer()
	if errLoadStepsTransfer != nil {
		return errLoadStepsTransfer
	}

	errCheckArgs := p.checkArgs()
	if errCheckArgs != nil {
		return fmt.Errorf("check args err: %v", errCheckArgs)
	}

	if p.onlyArgsCheck {
		wd_log.Info("only check args, skip do doBiz")
		return nil
	}

	err := p.doBiz()
	if err != nil {
		return err
	}
	errSaveStepsTransfer := p.saveStepsTransfer()
	if errSaveStepsTransfer != nil {
		return errSaveStepsTransfer
	}

	return nil
}

func (p *FileBrowserSyncPlugin) loadStepsTransfer() error {
	// change or remove or this code start
	//if p.Settings.StepsTransferDemo {
	//	var readConfigData Settings
	//	errLoad := wd_steps_transfer.In(p.Settings.RootPath, p.Settings.StepsTransferPath, *p.woodpeckerInfo, StepsTransferMarkDemoConfig, &readConfigData)
	//	if errLoad != nil {
	//		return nil
	//	}
	//	wd_log.VerboseJsonf(readConfigData, "load steps transfer config mark [ %s ]", StepsTransferMarkDemoConfig)
	//}
	// change or remove or this code end
	return nil
}

func (p *FileBrowserSyncPlugin) checkArgs() error {

	if p.Settings.SyncMode == "" {
		return fmt.Errorf("sync mode must set, now is empty, check flag [ %s ]", CliNameSyncMode)
	}

	errCheckSyncMode := argCheckInArr("sync mode", p.Settings.SyncMode, syncModeSupport)
	if errCheckSyncMode != nil {
		return errCheckSyncMode
	}

	if len(p.Settings.FileBrowserUrls) == 0 && p.Settings.FileBrowserStandbyUrl == "" {
		return fmt.Errorf("file browser urls and standby url all empty, please check")
	}

	if len(p.Settings.FileBrowserUrls) > 0 {
		if p.Settings.FileBrowserUsername == "" {
			return fmt.Errorf("file browser username must set, now is empty, check flag [ %s ]", CliNameFileBrowserUsername)
		}
		if p.Settings.FileBrowserUserPassword == "" {
			return fmt.Errorf("file browser user password must set, now is empty, check flag [ %s ]", CliNameFileBrowserUserPassword)
		}
	}
	if p.Settings.FileBrowserStandbyUrl != "" {
		if p.Settings.FileBrowserStandbyUsername == "" {
			return fmt.Errorf("file browser standby username must set, now is empty, check flag [ %s ]", CliNameFileBrowserStandbyUsername)
		}
		if p.Settings.FileBrowserStandbyUserPassword == "" {
			return fmt.Errorf("file browser standby user password must set, now is empty, check flag [ %s ]", CliNameFileBrowserStandbyUserPassword)
		}
	}

	if p.Settings.SyncWorkSpaceAbsPath == "" {
		return fmt.Errorf("sync work space path must set, now is empty, check flag [ %s ]", CliNameSyncWorkSpacePath)
	}

	if len(p.Settings.SyncIncludeGlobs) > 0 && len(p.Settings.SyncExcludeGlobs) > 0 {
		return fmt.Errorf("can not set include and exclude globs both, please remove one")
	}

	p.Settings.syncWorkSpacePath = filepath.Join(p.Settings.RootPath, p.Settings.SyncWorkSpaceAbsPath)

	// set default TimeoutSecond
	if p.Settings.TimeoutSecond < TimeoutSecondMinimum {
		p.Settings.TimeoutSecond = TimeoutSecondMinimum
	}
	// set default SyncTimeoutSecond
	if p.Settings.SyncTimeoutSecond < SyncTimeoutSecondMinimum {
		p.Settings.SyncTimeoutSecond = SyncTimeoutSecondMinimum
	}

	return nil
}

func argCheckInArr(mark string, target string, checkArr []string) error {
	if !(string_tools.StringInArr(target, checkArr)) {
		return fmt.Errorf("not support %s now [ %s ], must in %v", mark, target, checkArr)
	}
	return nil
}

// doBiz
//
//	replace this code with your file_browser_sync implementation
func (p *FileBrowserSyncPlugin) doBiz() error {

	p.chooseFileBrowserConnect()

	errSyncCheck := p.doSyncCheck()
	if errSyncCheck != nil {
		return errSyncCheck
	}

	errSyncByFileBrowser := p.doSyncByFileBrowser()
	if errSyncByFileBrowser != nil {
		return errSyncByFileBrowser
	}

	return nil
}

func (p *FileBrowserSyncPlugin) saveStepsTransfer() error {
	// change or remove this code

	if p.Settings.StepsOutDisable {
		wd_log.Debugf("steps out disable by flag [ %v ], skip save steps transfer", p.Settings.StepsOutDisable)
		return nil
	}

	// change or remove or this code start
	//if p.Settings.StepsTransferDemo {
	//	transferAppendObj, errSave := wd_steps_transfer.Out(p.Settings.RootPath, p.Settings.StepsTransferPath, *p.woodpeckerInfo, StepsTransferMarkDemoConfig, p.Settings)
	//	if errSave != nil {
	//		return errSave
	//	}
	//	wd_log.VerboseJsonf(transferAppendObj, "save steps transfer config mark [ %s ]", StepsTransferMarkDemoConfig)
	//}
	// change or remove or this code end
	return nil
}
