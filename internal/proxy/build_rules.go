package proxy

import (
	"github.com/universe-toolkits/chaosgate/internal/config"
	"github.com/universe-toolkits/chaosgate/internal/rules"
)

func buildRules(cfg *config.Config) []rules.Rule {
	var result []rules.Rule

	for _, rc := range cfg.Rules {
		var matchers []rules.Matcher
		matchers = append(matchers, rules.DefaultMatcher{
			Host:       rc.Match.Host,
			PathPrefix: rc.Match.PathPrefix,
			Method:     rc.Match.Method,
		})

		var action rules.Action

		switch rc.Action.Type {

		case "mock":
			action = &rules.MockAction{
				Status:  rc.Action.Status,
				Body:    rc.Action.Body,
				Headers: rc.Action.Header,
			}

		case "forward":
			action = &rules.ForwardAction{
				Target: rc.Action.Target,
			}

		case "drop":
			action = &rules.DropAction{}
		}

		result = append(result, rules.Rule{
			Name:       rc.Name,
			Percentage: rc.Percentage,
			Matchers:   matchers,
			Action:     action,
		})
	}

	return result
}
