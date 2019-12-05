package resource

type MemberResource struct {
	Project string
	Role string
	Member string
}

func NewMemberResource() *MemberResource {
	return &MemberResource{}
}
