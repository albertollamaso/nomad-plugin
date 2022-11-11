package nomad

import (
	"fmt"
	"strings"

	"github.com/hashicorp/nomad/api"
)

func getAllocImages(evt *api.Event) ([]string, error) {
	alloc, err := evt.Allocation()
	if err != nil {
		return nil, err
	}

	var images []string
	for _, task := range alloc.TaskStates {
		for _, taskEvent := range task.Events {
			image := taskEvent.Details["image"]
			fmt.Println("------------------")
			fmt.Println(image)
			fmt.Println("------------------")
			tokens := strings.Split(image, ":")
			if strings.Contains(tokens[0], "/") {
				// cases
				// [registry]/[repository_name]/[repo_path_component]/[image]:[tag] (repository name could have two or more path components, they must be separated by a forward slash (“/”).)
				// [registry]/[repository_name]/[image]:[tag]
				// [registry]/[image]:[tag]
				tokens = strings.Split(tokens[0], "/")
				images = append(images, tokens[len(tokens)-1])
			} else {
				// case [image]:[tag]
				images = append(images, tokens[0])
			}
		}
	}

	return images, nil
}

func getAllocTags(evt *api.Event) ([]string, error) {
	alloc, err := evt.Allocation()
	if err != nil {
		return nil, err
	}

	var tags []string
	for _, task := range alloc.TaskStates {
		for _, taskEvent := range task.Events {
			tokens := strings.Split(taskEvent.Details["image"], ":")
			tags = append(tags, tokens[len(tokens)-1])
		}
	}

	return tags, nil
}

func getAllocRepos(evt *api.Event) ([]string, error) {
	alloc, err := evt.Allocation()
	if err != nil {
		return nil, err
	}

	var repos []string
	for _, task := range alloc.TaskStates {
		for _, taskEvent := range task.Events {
			tokens := strings.Split(taskEvent.Details["image"], ":")
			if len(tokens) == 2 {
				if strings.Contains(tokens[0], "/") {
					// cases
					// [registry]/[repository_name]/[repo_path_component]/[image]:[tag] (repository name could have two or more path components, they must be separated by a forward slash (“/”).)
					// [registry]/[repository_name]/[image]:[tag]
					// [registry]/[image]:[tag]
					tokens = strings.Split(tokens[0], "/")
					tokens = tokens[:len(tokens)-1]
					tokenStr := strings.Join(tokens, "/")
					repos = append(repos, tokenStr)
				} else {
					// case [image]:[tag]
					repos = append(repos, "docker.io")
				}
			}
		}
	}

	return repos, nil
}
