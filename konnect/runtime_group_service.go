package konnect

import (
	"context"
	"encoding/json"
)

type RuntimeGroupService service

// List fetches a list of Service packages.
func (s *RuntimeGroupService) List(ctx context.Context,
	opt *ListOpt,
) ([]*RuntimeGroup, *ListOpt, error) {
	data, next, err := s.client.list(ctx, "/konnect-api/api/runtime_groups", opt)
	if err != nil {
		return nil, nil, err
	}
	var runtimeGroups []*RuntimeGroup

	for _, object := range data {
		b, err := object.MarshalJSON()
		if err != nil {
			return nil, nil, err
		}
		var runtimeGroup RuntimeGroup
		err = json.Unmarshal(b, &runtimeGroup)
		if err != nil {
			return nil, nil, err
		}
		runtimeGroups = append(runtimeGroups, &runtimeGroup)
	}

	return runtimeGroups, next, nil
}
