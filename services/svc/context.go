// GENERATED CODE - DO NOT EDIT!
// @generated
//
// Generated by:
//
//   go run gen_context.go
//
// Called via:
//
//   go generate
//

package svc

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sourcegraph.com/sourcegraph/sourcegraph/go-sourcegraph/sourcegraph"
	"sourcegraph.com/sourcegraph/srclib/store/pb"
)

type contextKey int

const (
	_MultiRepoImporterKey contextKey = iota
	_AccountsKey          contextKey = iota
	_AnnotationsKey       contextKey = iota
	_AuthKey              contextKey = iota
	_BuildsKey            contextKey = iota
	_DefsKey              contextKey = iota
	_DeltasKey            contextKey = iota
	_MetaKey              contextKey = iota
	_MirrorReposKey       contextKey = iota
	_NotifyKey            contextKey = iota
	_OrgsKey              contextKey = iota
	_PeopleKey            contextKey = iota
	_RepoStatusesKey      contextKey = iota
	_RepoTreeKey          contextKey = iota
	_ReposKey             contextKey = iota
	_UsersKey             contextKey = iota
)

// Services contains fields for all existing services.
type Services struct {
	MultiRepoImporter pb.MultiRepoImporterServer
	Accounts          sourcegraph.AccountsServer
	Annotations       sourcegraph.AnnotationsServer
	Auth              sourcegraph.AuthServer
	Builds            sourcegraph.BuildsServer
	Defs              sourcegraph.DefsServer
	Deltas            sourcegraph.DeltasServer
	Meta              sourcegraph.MetaServer
	MirrorRepos       sourcegraph.MirrorReposServer
	Notify            sourcegraph.NotifyServer
	Orgs              sourcegraph.OrgsServer
	People            sourcegraph.PeopleServer
	RepoStatuses      sourcegraph.RepoStatusesServer
	RepoTree          sourcegraph.RepoTreeServer
	Repos             sourcegraph.ReposServer
	Users             sourcegraph.UsersServer
}

// RegisterAll calls all of the the RegisterXxxServer funcs.
func RegisterAll(s *grpc.Server, svcs Services) {

	if svcs.MultiRepoImporter != nil {
		pb.RegisterMultiRepoImporterServer(s, svcs.MultiRepoImporter)
	}

	if svcs.Accounts != nil {
		sourcegraph.RegisterAccountsServer(s, svcs.Accounts)
	}

	if svcs.Annotations != nil {
		sourcegraph.RegisterAnnotationsServer(s, svcs.Annotations)
	}

	if svcs.Auth != nil {
		sourcegraph.RegisterAuthServer(s, svcs.Auth)
	}

	if svcs.Builds != nil {
		sourcegraph.RegisterBuildsServer(s, svcs.Builds)
	}

	if svcs.Defs != nil {
		sourcegraph.RegisterDefsServer(s, svcs.Defs)
	}

	if svcs.Deltas != nil {
		sourcegraph.RegisterDeltasServer(s, svcs.Deltas)
	}

	if svcs.Meta != nil {
		sourcegraph.RegisterMetaServer(s, svcs.Meta)
	}

	if svcs.MirrorRepos != nil {
		sourcegraph.RegisterMirrorReposServer(s, svcs.MirrorRepos)
	}

	if svcs.Notify != nil {
		sourcegraph.RegisterNotifyServer(s, svcs.Notify)
	}

	if svcs.Orgs != nil {
		sourcegraph.RegisterOrgsServer(s, svcs.Orgs)
	}

	if svcs.People != nil {
		sourcegraph.RegisterPeopleServer(s, svcs.People)
	}

	if svcs.RepoStatuses != nil {
		sourcegraph.RegisterRepoStatusesServer(s, svcs.RepoStatuses)
	}

	if svcs.RepoTree != nil {
		sourcegraph.RegisterRepoTreeServer(s, svcs.RepoTree)
	}

	if svcs.Repos != nil {
		sourcegraph.RegisterReposServer(s, svcs.Repos)
	}

	if svcs.Users != nil {
		sourcegraph.RegisterUsersServer(s, svcs.Users)
	}

}

// WithServices returns a copy of parent with the given services. If a service's field value is nil, its previous value is inherited from parent in the new context.
func WithServices(ctx context.Context, s Services) context.Context {

	if s.MultiRepoImporter != nil {
		ctx = WithMultiRepoImporter(ctx, s.MultiRepoImporter)
	}

	if s.Accounts != nil {
		ctx = WithAccounts(ctx, s.Accounts)
	}

	if s.Annotations != nil {
		ctx = WithAnnotations(ctx, s.Annotations)
	}

	if s.Auth != nil {
		ctx = WithAuth(ctx, s.Auth)
	}

	if s.Builds != nil {
		ctx = WithBuilds(ctx, s.Builds)
	}

	if s.Defs != nil {
		ctx = WithDefs(ctx, s.Defs)
	}

	if s.Deltas != nil {
		ctx = WithDeltas(ctx, s.Deltas)
	}

	if s.Meta != nil {
		ctx = WithMeta(ctx, s.Meta)
	}

	if s.MirrorRepos != nil {
		ctx = WithMirrorRepos(ctx, s.MirrorRepos)
	}

	if s.Notify != nil {
		ctx = WithNotify(ctx, s.Notify)
	}

	if s.Orgs != nil {
		ctx = WithOrgs(ctx, s.Orgs)
	}

	if s.People != nil {
		ctx = WithPeople(ctx, s.People)
	}

	if s.RepoStatuses != nil {
		ctx = WithRepoStatuses(ctx, s.RepoStatuses)
	}

	if s.RepoTree != nil {
		ctx = WithRepoTree(ctx, s.RepoTree)
	}

	if s.Repos != nil {
		ctx = WithRepos(ctx, s.Repos)
	}

	if s.Users != nil {
		ctx = WithUsers(ctx, s.Users)
	}

	return ctx
}

// WithMultiRepoImporter returns a copy of parent that uses the given MultiRepoImporter service.
func WithMultiRepoImporter(ctx context.Context, s pb.MultiRepoImporterServer) context.Context {
	return context.WithValue(ctx, _MultiRepoImporterKey, s)
}

// MultiRepoImporter gets the context's MultiRepoImporter service. If the service is not present, it panics.
func MultiRepoImporter(ctx context.Context) pb.MultiRepoImporterServer {
	s, ok := ctx.Value(_MultiRepoImporterKey).(pb.MultiRepoImporterServer)
	if !ok || s == nil {
		panic("no MultiRepoImporter set in context")
	}
	return s
}

// MultiRepoImporterOrNil returns the context's MultiRepoImporter service if present, or else nil.
func MultiRepoImporterOrNil(ctx context.Context) pb.MultiRepoImporterServer {
	s, ok := ctx.Value(_MultiRepoImporterKey).(pb.MultiRepoImporterServer)
	if ok {
		return s
	}
	return nil
}

// WithAccounts returns a copy of parent that uses the given Accounts service.
func WithAccounts(ctx context.Context, s sourcegraph.AccountsServer) context.Context {
	return context.WithValue(ctx, _AccountsKey, s)
}

// Accounts gets the context's Accounts service. If the service is not present, it panics.
func Accounts(ctx context.Context) sourcegraph.AccountsServer {
	s, ok := ctx.Value(_AccountsKey).(sourcegraph.AccountsServer)
	if !ok || s == nil {
		panic("no Accounts set in context")
	}
	return s
}

// AccountsOrNil returns the context's Accounts service if present, or else nil.
func AccountsOrNil(ctx context.Context) sourcegraph.AccountsServer {
	s, ok := ctx.Value(_AccountsKey).(sourcegraph.AccountsServer)
	if ok {
		return s
	}
	return nil
}

// WithAnnotations returns a copy of parent that uses the given Annotations service.
func WithAnnotations(ctx context.Context, s sourcegraph.AnnotationsServer) context.Context {
	return context.WithValue(ctx, _AnnotationsKey, s)
}

// Annotations gets the context's Annotations service. If the service is not present, it panics.
func Annotations(ctx context.Context) sourcegraph.AnnotationsServer {
	s, ok := ctx.Value(_AnnotationsKey).(sourcegraph.AnnotationsServer)
	if !ok || s == nil {
		panic("no Annotations set in context")
	}
	return s
}

// AnnotationsOrNil returns the context's Annotations service if present, or else nil.
func AnnotationsOrNil(ctx context.Context) sourcegraph.AnnotationsServer {
	s, ok := ctx.Value(_AnnotationsKey).(sourcegraph.AnnotationsServer)
	if ok {
		return s
	}
	return nil
}

// WithAuth returns a copy of parent that uses the given Auth service.
func WithAuth(ctx context.Context, s sourcegraph.AuthServer) context.Context {
	return context.WithValue(ctx, _AuthKey, s)
}

// Auth gets the context's Auth service. If the service is not present, it panics.
func Auth(ctx context.Context) sourcegraph.AuthServer {
	s, ok := ctx.Value(_AuthKey).(sourcegraph.AuthServer)
	if !ok || s == nil {
		panic("no Auth set in context")
	}
	return s
}

// AuthOrNil returns the context's Auth service if present, or else nil.
func AuthOrNil(ctx context.Context) sourcegraph.AuthServer {
	s, ok := ctx.Value(_AuthKey).(sourcegraph.AuthServer)
	if ok {
		return s
	}
	return nil
}

// WithBuilds returns a copy of parent that uses the given Builds service.
func WithBuilds(ctx context.Context, s sourcegraph.BuildsServer) context.Context {
	return context.WithValue(ctx, _BuildsKey, s)
}

// Builds gets the context's Builds service. If the service is not present, it panics.
func Builds(ctx context.Context) sourcegraph.BuildsServer {
	s, ok := ctx.Value(_BuildsKey).(sourcegraph.BuildsServer)
	if !ok || s == nil {
		panic("no Builds set in context")
	}
	return s
}

// BuildsOrNil returns the context's Builds service if present, or else nil.
func BuildsOrNil(ctx context.Context) sourcegraph.BuildsServer {
	s, ok := ctx.Value(_BuildsKey).(sourcegraph.BuildsServer)
	if ok {
		return s
	}
	return nil
}

// WithDefs returns a copy of parent that uses the given Defs service.
func WithDefs(ctx context.Context, s sourcegraph.DefsServer) context.Context {
	return context.WithValue(ctx, _DefsKey, s)
}

// Defs gets the context's Defs service. If the service is not present, it panics.
func Defs(ctx context.Context) sourcegraph.DefsServer {
	s, ok := ctx.Value(_DefsKey).(sourcegraph.DefsServer)
	if !ok || s == nil {
		panic("no Defs set in context")
	}
	return s
}

// DefsOrNil returns the context's Defs service if present, or else nil.
func DefsOrNil(ctx context.Context) sourcegraph.DefsServer {
	s, ok := ctx.Value(_DefsKey).(sourcegraph.DefsServer)
	if ok {
		return s
	}
	return nil
}

// WithDeltas returns a copy of parent that uses the given Deltas service.
func WithDeltas(ctx context.Context, s sourcegraph.DeltasServer) context.Context {
	return context.WithValue(ctx, _DeltasKey, s)
}

// Deltas gets the context's Deltas service. If the service is not present, it panics.
func Deltas(ctx context.Context) sourcegraph.DeltasServer {
	s, ok := ctx.Value(_DeltasKey).(sourcegraph.DeltasServer)
	if !ok || s == nil {
		panic("no Deltas set in context")
	}
	return s
}

// DeltasOrNil returns the context's Deltas service if present, or else nil.
func DeltasOrNil(ctx context.Context) sourcegraph.DeltasServer {
	s, ok := ctx.Value(_DeltasKey).(sourcegraph.DeltasServer)
	if ok {
		return s
	}
	return nil
}

// WithMeta returns a copy of parent that uses the given Meta service.
func WithMeta(ctx context.Context, s sourcegraph.MetaServer) context.Context {
	return context.WithValue(ctx, _MetaKey, s)
}

// Meta gets the context's Meta service. If the service is not present, it panics.
func Meta(ctx context.Context) sourcegraph.MetaServer {
	s, ok := ctx.Value(_MetaKey).(sourcegraph.MetaServer)
	if !ok || s == nil {
		panic("no Meta set in context")
	}
	return s
}

// MetaOrNil returns the context's Meta service if present, or else nil.
func MetaOrNil(ctx context.Context) sourcegraph.MetaServer {
	s, ok := ctx.Value(_MetaKey).(sourcegraph.MetaServer)
	if ok {
		return s
	}
	return nil
}

// WithMirrorRepos returns a copy of parent that uses the given MirrorRepos service.
func WithMirrorRepos(ctx context.Context, s sourcegraph.MirrorReposServer) context.Context {
	return context.WithValue(ctx, _MirrorReposKey, s)
}

// MirrorRepos gets the context's MirrorRepos service. If the service is not present, it panics.
func MirrorRepos(ctx context.Context) sourcegraph.MirrorReposServer {
	s, ok := ctx.Value(_MirrorReposKey).(sourcegraph.MirrorReposServer)
	if !ok || s == nil {
		panic("no MirrorRepos set in context")
	}
	return s
}

// MirrorReposOrNil returns the context's MirrorRepos service if present, or else nil.
func MirrorReposOrNil(ctx context.Context) sourcegraph.MirrorReposServer {
	s, ok := ctx.Value(_MirrorReposKey).(sourcegraph.MirrorReposServer)
	if ok {
		return s
	}
	return nil
}

// WithNotify returns a copy of parent that uses the given Notify service.
func WithNotify(ctx context.Context, s sourcegraph.NotifyServer) context.Context {
	return context.WithValue(ctx, _NotifyKey, s)
}

// Notify gets the context's Notify service. If the service is not present, it panics.
func Notify(ctx context.Context) sourcegraph.NotifyServer {
	s, ok := ctx.Value(_NotifyKey).(sourcegraph.NotifyServer)
	if !ok || s == nil {
		panic("no Notify set in context")
	}
	return s
}

// NotifyOrNil returns the context's Notify service if present, or else nil.
func NotifyOrNil(ctx context.Context) sourcegraph.NotifyServer {
	s, ok := ctx.Value(_NotifyKey).(sourcegraph.NotifyServer)
	if ok {
		return s
	}
	return nil
}

// WithOrgs returns a copy of parent that uses the given Orgs service.
func WithOrgs(ctx context.Context, s sourcegraph.OrgsServer) context.Context {
	return context.WithValue(ctx, _OrgsKey, s)
}

// Orgs gets the context's Orgs service. If the service is not present, it panics.
func Orgs(ctx context.Context) sourcegraph.OrgsServer {
	s, ok := ctx.Value(_OrgsKey).(sourcegraph.OrgsServer)
	if !ok || s == nil {
		panic("no Orgs set in context")
	}
	return s
}

// OrgsOrNil returns the context's Orgs service if present, or else nil.
func OrgsOrNil(ctx context.Context) sourcegraph.OrgsServer {
	s, ok := ctx.Value(_OrgsKey).(sourcegraph.OrgsServer)
	if ok {
		return s
	}
	return nil
}

// WithPeople returns a copy of parent that uses the given People service.
func WithPeople(ctx context.Context, s sourcegraph.PeopleServer) context.Context {
	return context.WithValue(ctx, _PeopleKey, s)
}

// People gets the context's People service. If the service is not present, it panics.
func People(ctx context.Context) sourcegraph.PeopleServer {
	s, ok := ctx.Value(_PeopleKey).(sourcegraph.PeopleServer)
	if !ok || s == nil {
		panic("no People set in context")
	}
	return s
}

// PeopleOrNil returns the context's People service if present, or else nil.
func PeopleOrNil(ctx context.Context) sourcegraph.PeopleServer {
	s, ok := ctx.Value(_PeopleKey).(sourcegraph.PeopleServer)
	if ok {
		return s
	}
	return nil
}

// WithRepoStatuses returns a copy of parent that uses the given RepoStatuses service.
func WithRepoStatuses(ctx context.Context, s sourcegraph.RepoStatusesServer) context.Context {
	return context.WithValue(ctx, _RepoStatusesKey, s)
}

// RepoStatuses gets the context's RepoStatuses service. If the service is not present, it panics.
func RepoStatuses(ctx context.Context) sourcegraph.RepoStatusesServer {
	s, ok := ctx.Value(_RepoStatusesKey).(sourcegraph.RepoStatusesServer)
	if !ok || s == nil {
		panic("no RepoStatuses set in context")
	}
	return s
}

// RepoStatusesOrNil returns the context's RepoStatuses service if present, or else nil.
func RepoStatusesOrNil(ctx context.Context) sourcegraph.RepoStatusesServer {
	s, ok := ctx.Value(_RepoStatusesKey).(sourcegraph.RepoStatusesServer)
	if ok {
		return s
	}
	return nil
}

// WithRepoTree returns a copy of parent that uses the given RepoTree service.
func WithRepoTree(ctx context.Context, s sourcegraph.RepoTreeServer) context.Context {
	return context.WithValue(ctx, _RepoTreeKey, s)
}

// RepoTree gets the context's RepoTree service. If the service is not present, it panics.
func RepoTree(ctx context.Context) sourcegraph.RepoTreeServer {
	s, ok := ctx.Value(_RepoTreeKey).(sourcegraph.RepoTreeServer)
	if !ok || s == nil {
		panic("no RepoTree set in context")
	}
	return s
}

// RepoTreeOrNil returns the context's RepoTree service if present, or else nil.
func RepoTreeOrNil(ctx context.Context) sourcegraph.RepoTreeServer {
	s, ok := ctx.Value(_RepoTreeKey).(sourcegraph.RepoTreeServer)
	if ok {
		return s
	}
	return nil
}

// WithRepos returns a copy of parent that uses the given Repos service.
func WithRepos(ctx context.Context, s sourcegraph.ReposServer) context.Context {
	return context.WithValue(ctx, _ReposKey, s)
}

// Repos gets the context's Repos service. If the service is not present, it panics.
func Repos(ctx context.Context) sourcegraph.ReposServer {
	s, ok := ctx.Value(_ReposKey).(sourcegraph.ReposServer)
	if !ok || s == nil {
		panic("no Repos set in context")
	}
	return s
}

// ReposOrNil returns the context's Repos service if present, or else nil.
func ReposOrNil(ctx context.Context) sourcegraph.ReposServer {
	s, ok := ctx.Value(_ReposKey).(sourcegraph.ReposServer)
	if ok {
		return s
	}
	return nil
}

// WithUsers returns a copy of parent that uses the given Users service.
func WithUsers(ctx context.Context, s sourcegraph.UsersServer) context.Context {
	return context.WithValue(ctx, _UsersKey, s)
}

// Users gets the context's Users service. If the service is not present, it panics.
func Users(ctx context.Context) sourcegraph.UsersServer {
	s, ok := ctx.Value(_UsersKey).(sourcegraph.UsersServer)
	if !ok || s == nil {
		panic("no Users set in context")
	}
	return s
}

// UsersOrNil returns the context's Users service if present, or else nil.
func UsersOrNil(ctx context.Context) sourcegraph.UsersServer {
	s, ok := ctx.Value(_UsersKey).(sourcegraph.UsersServer)
	if ok {
		return s
	}
	return nil
}
