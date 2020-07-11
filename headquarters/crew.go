package main

import (
	"context"
	"net/http"

	pb "github.com/ides15/planet-express/ship/pkg/planetexpress"
)

type crewResolver struct {
	c pb.Crew
}

func (r *resolver) Crew(ctx context.Context) (*crewResolver, error) {
	crew, err := getCrew(client)
	// With more time I'd add better error handling of this, since this error might not always be not found
	if err != nil {
		return nil, notFoundError{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}
	}

	return &crewResolver{
		c: crew,
	}, nil
}

func (c *crewResolver) Size(ctx context.Context) int32 {
	return int32(c.c.GetSize())
}

func (c *crewResolver) Mood(ctx context.Context) pb.Crew_Mood {
	return c.c.GetMood()
}

func (c *crewResolver) CaptainMarooned(ctx context.Context) bool {
	return c.c.GetCaptainMarooned()
}

type memberResolver struct {
	m pb.Member
}

func (m *memberResolver) ID(ctx context.Context) string {
	return m.m.GetId()
}

func (m *memberResolver) FirstName(ctx context.Context) string {
	return m.m.GetFirstName()
}

func (m *memberResolver) LastName(ctx context.Context) string {
	return m.m.GetLastName()
}

func (m *memberResolver) Age(ctx context.Context) int32 {
	return int32(m.m.GetAge())
}

func (c *crewResolver) Members(ctx context.Context) *[]*memberResolver {
	members := c.c.GetMembers()

	mr := make([]*memberResolver, len(members))
	for i := range members {
		mr[i] = &memberResolver{
			m: *members[i],
		}
	}

	return &mr
}
