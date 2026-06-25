package profile

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

func Create(name, user, project string) error {
	if err := validateName(name); err != nil {
		return err
	}

	filename := fileName(name)

	if _, err := os.Stat(filename); err == nil {
		return fmt.Errorf("profile %q already exists", name)
	}

	p := Profile{
		User:    user,
		Project: project,
	}

	data, err := yaml.Marshal(&p)
	if err != nil {
		return fmt.Errorf("marshal yaml: %w", err)
	}

	if err := os.WriteFile(filename, data, 0644); err != nil {
		return fmt.Errorf("write file: %w", err)
	}

	return nil
}

func Get(name string) (Profile, error) {
	if err := validateName(name); err != nil {
		return Profile{}, err
	}

	data, err := os.ReadFile(fileName(name))
	if err != nil {
		if os.IsNotExist(err) {
			return Profile{}, fmt.Errorf("profile %q not found", name)
		}
		return Profile{}, fmt.Errorf("read file: %w", err)
	}

	var p Profile
	if err := yaml.Unmarshal(data, &p); err != nil {
		return Profile{}, fmt.Errorf("unmarshal yaml: %w", err)
	}

	return p, nil
}

func List() ([]NamedProfile, error) {
	entries, err := os.ReadDir(".")
	if err != nil {
		return nil, fmt.Errorf("read dir: %w", err)
	}

	var result []NamedProfile

	for _, e := range entries {
		if e.IsDir() {
			continue
		}

		name := e.Name()
		if !strings.HasSuffix(name, ".yaml") && !strings.HasSuffix(name, ".yml") {
			continue
		}

		data, err := os.ReadFile(name)
		if err != nil {
			return nil, fmt.Errorf("read file %s: %w", name, err)
		}

		var p Profile
		if err := yaml.Unmarshal(data, &p); err != nil {
			return nil, fmt.Errorf("unmarshal yaml %s: %w", name, err)
		}

		profileName := strings.TrimSuffix(strings.TrimSuffix(name, ".yaml"), ".yml")

		result = append(result, NamedProfile{
			Name:    profileName,
			User:    p.User,
			Project: p.Project,
		})
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name < result[j].Name
	})

	return result, nil
}

func Delete(name string) error {
	if err := validateName(name); err != nil {
		return err
	}

	err := os.Remove(fileName(name))
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("profile %q not found", name)
		}
		return fmt.Errorf("delete file: %w", err)
	}

	return nil
}

func fileName(name string) string {
	return name + ".yaml"
}

func validateName(name string) error {
	if name == "" {
		return fmt.Errorf("empty profile name")
	}

	for _, r := range name {
		if (r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') ||
			r == '-' || r == '_' {
			continue
		}
		return fmt.Errorf("invalid profile name %q: only letters, digits, '-' and '_' are allowed", name)
	}

	return nil
}
