{{- $alias := .Aliases.Table .Table.Name -}}

type Create{{$alias.UpSingular}}Payload struct {
	Items []*{{$alias.UpSingular}}
}

func (p Create{{$alias.UpSingular}}Payload) Item() *{{$alias.UpSingular}} {
	if len(p.Items) > 0 {
		return p.Items[0]
	}
	return nil
}

func (p Create{{$alias.UpSingular}}Payload) ID() *uuid.UUID {
	item := p.Item()
	if item == nil {
		return nil
	}
	return &item.ID
}

func (r *Repository) Create{{$alias.UpPlural}}(ctx context.Context, input []*{{$alias.UpSingular}}) (*Create{{$alias.UpSingular}}Payload, error) {
	res := make([]*{{$alias.UpSingular}}, 0, len(input))
	for _, m := range input {
		if m.ID == uuid.Nil {
			var err error
			m.ID, err = uuid.NewV4()
			if err != nil {
				return nil, err
			}
		}
		err := m.Insert(ctx, r.DB(ctx), boil.Infer())
		if err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return &Create{{$alias.UpSingular}}Payload{Items: res}, nil
}

type Update{{$alias.UpSingular}}Payload struct {
	Items []*{{$alias.UpSingular}}
}

func (p Update{{$alias.UpSingular}}Payload) Item() *{{$alias.UpSingular}} {
	if len(p.Items) > 0 {
		return p.Items[0]
	}
	return nil
}

func (r *Repository) Update{{$alias.UpPlural}}(ctx context.Context, input []*{{$alias.UpSingular}}) (*Update{{$alias.UpSingular}}Payload, error) {
	res := make([]*{{$alias.UpSingular}}, 0, len(input))
	for _, m := range input {
		_, err := m.Update(ctx, r.DB(ctx), boil.Infer())

		if err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return &Update{{$alias.UpSingular}}Payload{Items: res}, nil
}


type Delete{{$alias.UpSingular}}Payload struct {
	IDs []uuid.UUID
}

func (p Delete{{$alias.UpSingular}}Payload) ID() *uuid.UUID {
	if len(p.IDs) > 0 {
		return &p.IDs[0]
	}
	return nil
}

func (r *Repository) Delete{{$alias.UpPlural}}(ctx context.Context, input []uuid.UUID) (*Delete{{$alias.UpSingular}}Payload, error) {
	res := make([]uuid.UUID, 0, len(input))
	for _, id := range input {
		_, err := (&{{$alias.UpSingular}}{ID: id}).Delete(ctx, r.DB(ctx))

		if err != nil {
			return nil, err
		}
		res = append(res, id)
	}
	return &Delete{{$alias.UpSingular}}Payload{IDs: res}, nil
}

func (r *Repository) {{$alias.UpPlural}}(ctx context.Context, filter {{$alias.UpSingular}}Filter) ([]*{{$alias.UpSingular}}, error) {
	// TODO Apply filters
	return {{$alias.UpPlural}}().All(ctx, r.DB(ctx))
}
