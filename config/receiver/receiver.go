// Copyright 2023 Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package receiver

import (
	"github.com/go-kit/log"

	commoncfg "github.com/prometheus/common/config"

	"github.com/tyr1k/alertmanager/config"
	"github.com/tyr1k/alertmanager/notify"
	"github.com/tyr1k/alertmanager/notify/discord"
	"github.com/tyr1k/alertmanager/notify/email"
	"github.com/tyr1k/alertmanager/notify/msteams"
	"github.com/tyr1k/alertmanager/notify/opsgenie"
	"github.com/tyr1k/alertmanager/notify/pagerduty"
	"github.com/tyr1k/alertmanager/notify/pushover"
	"github.com/tyr1k/alertmanager/notify/slack"
	"github.com/tyr1k/alertmanager/notify/sns"
	"github.com/tyr1k/alertmanager/notify/telegram"
	"github.com/tyr1k/alertmanager/notify/victorops"
	"github.com/tyr1k/alertmanager/notify/webex"
	"github.com/tyr1k/alertmanager/notify/webhook"
	"github.com/tyr1k/alertmanager/notify/wechat"
	"github.com/tyr1k/alertmanager/template"
	"github.com/tyr1k/alertmanager/types"
)

// BuildReceiverIntegrations builds a list of integration notifiers off of a
// receiver config.
func BuildReceiverIntegrations(nc config.Receiver, tmpl *template.Template, logger log.Logger, httpOpts ...commoncfg.HTTPClientOption) ([]notify.Integration, error) {
	var (
		errs         types.MultiError
		integrations []notify.Integration
		add          = func(name string, i int, rs notify.ResolvedSender, f func(l log.Logger) (notify.Notifier, error)) {
			n, err := f(log.With(logger, "integration", name))
			if err != nil {
				errs.Add(err)
				return
			}
			integrations = append(integrations, notify.NewIntegration(n, rs, name, i, nc.Name))
		}
	)

	for i, c := range nc.WebhookConfigs {
		add("webhook", i, c, func(l log.Logger) (notify.Notifier, error) { return webhook.New(c, tmpl, l, httpOpts...) })
	}
	for i, c := range nc.EmailConfigs {
		add("email", i, c, func(l log.Logger) (notify.Notifier, error) { return email.New(c, tmpl, l), nil })
	}
	for i, c := range nc.PagerdutyConfigs {
		add("pagerduty", i, c, func(l log.Logger) (notify.Notifier, error) { return pagerduty.New(c, tmpl, l, httpOpts...) })
	}
	for i, c := range nc.OpsGenieConfigs {
		add("opsgenie", i, c, func(l log.Logger) (notify.Notifier, error) { return opsgenie.New(c, tmpl, l, httpOpts...) })
	}
	for i, c := range nc.WechatConfigs {
		add("wechat", i, c, func(l log.Logger) (notify.Notifier, error) { return wechat.New(c, tmpl, l, httpOpts...) })
	}
	for i, c := range nc.SlackConfigs {
		add("slack", i, c, func(l log.Logger) (notify.Notifier, error) { return slack.New(c, tmpl, l, httpOpts...) })
	}
	for i, c := range nc.VictorOpsConfigs {
		add("victorops", i, c, func(l log.Logger) (notify.Notifier, error) { return victorops.New(c, tmpl, l, httpOpts...) })
	}
	for i, c := range nc.PushoverConfigs {
		add("pushover", i, c, func(l log.Logger) (notify.Notifier, error) { return pushover.New(c, tmpl, l, httpOpts...) })
	}
	for i, c := range nc.SNSConfigs {
		add("sns", i, c, func(l log.Logger) (notify.Notifier, error) { return sns.New(c, tmpl, l, httpOpts...) })
	}
	for i, c := range nc.TelegramConfigs {
		add("telegram", i, c, func(l log.Logger) (notify.Notifier, error) { return telegram.New(c, tmpl, l, httpOpts...) })
	}
	for i, c := range nc.DiscordConfigs {
		add("discord", i, c, func(l log.Logger) (notify.Notifier, error) { return discord.New(c, tmpl, l, httpOpts...) })
	}
	for i, c := range nc.WebexConfigs {
		add("webex", i, c, func(l log.Logger) (notify.Notifier, error) { return webex.New(c, tmpl, l, httpOpts...) })
	}
	for i, c := range nc.MSTeamsConfigs {
		add("msteams", i, c, func(l log.Logger) (notify.Notifier, error) { return msteams.New(c, tmpl, l, httpOpts...) })
	}

	if errs.Len() > 0 {
		return nil, &errs
	}
	return integrations, nil
}
