package metallb

import "github.com/kubeclipper/kubeclipper/pkg/component"

func initI18nForComponentMeta() error {
	return component.AddI18nMessages(component.I18nMessages{
		{
			ID:      "metallb.metaTitle",
			English: "metallb Setting",
			Chinese: "metallb 设置",
		},
		{
			ID:      "metallb.mode",
			English: "Mode",
			Chinese: "模式",
		},
		{
			ID:      "metallb.addresses",
			English: "IPAddressPool",
			Chinese: "IP 地址池",
		},
		{
			ID:      "metallb.imageRepoMirror",
			English: "metallb Image Repository Mirror",
			Chinese: "metallb 镜像仓库代理",
		},
	})
}
