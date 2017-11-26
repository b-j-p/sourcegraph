// +build !dev

package graphqlbackend

// Code generated by schema_generate.go

// Schema is the raw graqhql schema
var Schema = `# Run this before committing changes to this file
# go generate sourcegraph.com/sourcegraph/sourcegraph/cmd/frontend/internal/graphqlbackend

schema {
  query: Query
  mutation: Mutation
}

type EmptyResponse {
  alwaysNil: String
}

interface Node {
  id: ID!
}

type Query {
  root: Root!
  node(id: ID!): Node
}

type ThreadLines {
  # HTML context lines before 'html'.
  #
  # It is sanitized already by the server, and thus is safe for rendering.
  htmlBefore: String!
  # HTML lines that the user's selection was made on.
  #
  # It is sanitized already by the server, and thus is safe for rendering.
  html: String!
  # HTML context lines after 'html'.
  #
  # It is sanitized already by the server, and thus is safe for rendering.
  htmlAfter: String!
  # text context lines before 'text'.
  textBefore: String!
  # text lines that the user's selection was made on.
  text: String!
  # text context lines after 'text'.
  textAfter: String!
  # byte offset into textLines where user selection began
  #
  # In Go syntax, userSelection := text[rangeStart:rangeStart+rangeLength]
  textSelectionRangeStart: Int!
  # length in bytes of the user selection
  #
  # In Go syntax, userSelection := text[rangeStart:rangeStart+rangeLength]
  textSelectionRangeLength: Int!
}

# Literally the exact same thing as above, except it's an input type because
# GraphQL doesn't allow mixing input and output types.
input ThreadLinesInput {
  # HTML context lines before 'html'.
  htmlBefore: String!
  # HTML lines that the user's selection was made on.
  html: String!
  # HTML context lines after 'html'.
  htmlAfter: String!
  # text context lines before 'text'.
  textBefore: String!
  # text lines that the user's selection was made on.
  text: String!
  # text context lines after 'text'.
  textAfter: String!
  # byte offset into textLines where user selection began
  #
  # In Go syntax, userSelection := text[rangeStart:rangeStart+rangeLength]
  textSelectionRangeStart: Int!
  # length in bytes of the user selection
  #
  # In Go syntax, userSelection := text[rangeStart:rangeStart+rangeLength]
  textSelectionRangeLength: Int!
}

type Mutation {
  createUser(username: String!, displayName: String!, avatarURL: String): User!
  createThread(
    # Temporarily until all clients are migrated, the orgInt32OrID type allows
    # orgID to be specified either as an ID! or Int!.
    orgID: ID!
    canonicalRemoteID: String!
    cloneURL: String!
    file: String!
    repoRevision: String!
    linesRevision: String!
    branch: String
    startLine: Int!
    endLine: Int!
    startCharacter: Int!
    endCharacter: Int!
    rangeLength: Int!
    contents: String!
    lines: ThreadLinesInput
  ): Thread!
  updateUser(username: String, displayName: String, avatarURL: String): User!
  # Update the settings for the currently authenticated user.
  updateUserSettings(lastKnownSettingsID: Int, contents: String!): Settings!
  updateThread(threadID: Int!, archived: Boolean): Thread!
  addCommentToThread(threadID: Int!, contents: String!): Thread!
  # This method is the same as addCommentToThread, the only difference is
  # that authentication is based on the secret ULID instead of the current
  # user.
  #
  # 🚨 SECURITY: Every field of the return type here is accessible publicly
  # given a shared item URL.
  addCommentToThreadShared(
    ulid: String!
    threadID: Int!
    contents: String!
  ): SharedItemThread!
  shareThread(threadID: Int!): String!
  shareComment(commentID: Int!): String!
  createOrg(name: String!, displayName: String!): Org!
  updateOrg(id: ID!, displayName: String, slackWebhookURL: String): Org!
  updateOrgSettings(
    # The ID of the org whose settings should be updated.
    id: ID
    # TODO(sqs): orgID is deprecated. Use id instead.
    orgID: ID
    lastKnownSettingsID: Int
    contents: String!
  ): Settings!
  inviteUser(email: String!, orgID: ID!): EmptyResponse
  acceptUserInvite(inviteToken: String!): OrgInviteStatus!
  removeUserFromOrg(userID: String!, orgID: ID!): EmptyResponse
  # adds a phabricator repository to the Sourcegraph server.
  # example callsign: "MUX"
  # example uri: "github.com/gorilla/mux"
  addPhabricatorRepo(callsign: String!, uri: String!): EmptyResponse
  logUserEvent(event: UserEvent!): EmptyResponse
}

type Root {
  repository(uri: String!): Repository
  phabricatorRepo(uri: String!): PhabricatorRepo
  repositories(query: String = ""): [Repository!]!
  symbols(id: String!, mode: String!): [Symbol!]!
  currentUser: User
  isUsernameAvailable(username: String!): Boolean!
  configuration: ConfigurationCascade!
  search2(query: String = "", scopeQuery: String = ""): Search2
  searchScopes2: [SearchScope2!]!
  repoGroups: [RepoGroup!]!
  revealCustomerCompany(ip: String!): CompanyProfile
  org(id: ID!): Org! @deprecated(reason: "use Query.node instead")
  sharedItem(ulid: String!): SharedItem
  packages(
    lang: String!
    id: String
    type: String
    name: String
    commit: String
    baseDir: String
    repoURL: String
    version: String
    limit: Int
  ): [Package!]!
  dependents(
    lang: String!
    id: String
    type: String
    name: String
    commit: String
    baseDir: String
    repoURL: String
    version: String
    package: String
    limit: Int
  ): [Dependency!]!
  users: [User!]!
}

type Search2 {
  results: SearchResults2!
  suggestions(first: Int): [SearchSuggestion2!]!
}

union SearchResult = FileMatch

type SearchResults2 {
  results: [SearchResult!]!
  limitHit: Boolean!
  cloning: [String!]!
  missing: [String!]!
  # An alert message that should be displayed before any results.
  alert: SearchAlert
}

union SearchSuggestion2 = Repository | File

type SearchScope2 {
  name: String!
  value: String!
}

# A search-related alert message.
type SearchAlert {
  title: String!
  description: String
  # "Did you mean: ____" query proposals
  proposedQueries: [SearchQuery2Description!]
}

type SearchQuery2Description {
  description: String
  query: SearchQuery2!
}

type SearchQuery2 {
  query: String!
  scopeQuery: String!
}

# A group of repositories.
type RepoGroup {
  name: String!
  repositories: [String!]!
}

# Represents a shared item (either a shared code comment OR code snippet).
#
# 🚨 SECURITY: Every field here is accessible publicly given a shared item URL.
# Do NOT use any non-primitive graphql type here unless it is also a SharedItem
# type.
type SharedItem {
  # who shared the item.
  author: SharedItemUser!
  public: Boolean!
  thread: SharedItemThread!
  # present only if the shared item was a specific comment.
  comment: SharedItemComment
}

# Like the User type, except with fields that should not be accessible with a
# secret URL removed.
#
# 🚨 SECURITY: Every field here is accessible publicly given a shared item URL.
# Do NOT use any non-primitive graphql type here unless it is also a SharedItem
# type.
type SharedItemUser {
  displayName: String
  username: String
  avatarURL: String
}

# Like the Thread type, except with fields that should not be accessible with a
# secret URL removed.
#
# 🚨 SECURITY: Every field here is accessible publicly given a shared item URL.
# Do NOT use any non-primitive graphql type here unless it is also a SharedItem
# type.
type SharedItemThread {
  id: Int!
  repo: SharedItemOrgRepo!
  file: String!
  branch: String
  repoRevision: String!
  title: String!
  startLine: Int!
  endLine: Int!
  startCharacter: Int!
  endCharacter: Int!
  rangeLength: Int!
  createdAt: String!
  archivedAt: String
  author: SharedItemUser!
  lines: SharedItemThreadLines
  comments: [SharedItemComment!]!
}

# Like the OrgRepo type, except with fields that should not be accessible with
# a secret URL removed.
#
# 🚨 SECURITY: Every field here is accessible publicly given a shared item URL.
# Do NOT use any non-primitive graphql type here unless it is also a SharedItem
# type.
type SharedItemOrgRepo {
  id: Int!
  remoteUri: String!
}

# Like the Comment type, except with fields that should not be accessible with a
# secret URL removed.
#
# 🚨 SECURITY: Every field here is accessible publicly given a shared item URL.
# Do NOT use any non-primitive graphql type here unless it is also a SharedItem
# type.
type SharedItemComment {
  id: Int!
  title: String!
  contents: String!
  createdAt: String!
  updatedAt: String!
  author: SharedItemUser!
}

# Exactly the same as the ThreadLines type, except it cannot have sensitive
# fields accidently added.
#
# 🚨 SECURITY: Every field here is accessible publicly given a shared item URL.
# Do NOT use any non-primitive graphql type here unless it is also a SharedItem
# type.
type SharedItemThreadLines {
  htmlBefore: String!
  html: String!
  htmlAfter: String!
  textBefore: String!
  text: String!
  textAfter: String!
  textSelectionRangeStart: Int!
  textSelectionRangeLength: Int!
}

type RefFields {
  refLocation: RefLocation
  uri: URI
}

type URI {
  host: String!
  fragment: String!
  path: String!
  query: String!
  scheme: String!
}

type RefLocation {
  startLineNumber: Int!
  startColumn: Int!
  endLineNumber: Int!
  endColumn: Int!
}

type Repository implements Node {
  id: ID!
  uri: String!
  description: String!
  language: String!
  fork: Boolean!
  starsCount: Int
  forksCount: Int
  private: Boolean!
  createdAt: String!
  pushedAt: String!
  commit(rev: String!): CommitState!
  revState(rev: String!): RevState!
  latest: CommitState!
  lastIndexedRevOrLatest: CommitState!
  defaultBranch: String!
  branches: [String!]!
  tags: [String!]!
  listTotalRefs: TotalRefList!
  gitCmdRaw(params: [String!]!): String!
  # Link to another sourcegraph instance location where this repository is located.
  redirectURL: String
}

type PhabricatorRepo {
  # the canonical repo path, like 'github.com/gorilla/mux'
  uri: String!
  # the unique Phabricator identifier for the repo, like 'MUX'
  callsign: String!
}

type TotalRefList {
  repositories: [Repository!]!
  total: Int!
}

type Symbol {
  repository: Repository!
  path: String!
  line: Int!
  character: Int!
}

type CommitState {
  commit: Commit
  cloneInProgress: Boolean!
}

type RevState {
  commit: Commit
  cloneInProgress: Boolean!
}

type Commit implements Node {
  id: ID!
  sha1: String!
  tree(path: String = "", recursive: Boolean = false): Tree
  file(path: String!): File
  languages: [String!]!
}

type CommitInfo {
  rev: String!
  author: Signature
  committer: Signature
  message: String!
}

type Signature {
  person: Person
  date: String!
}

type Person {
  name: String!
  email: String!
  gravatarHash: String!
  avatarURL: String!
}

type Tree {
  directories: [Directory]!
  files: [File]!
  # Consists of directories plus files.
  entries: [TreeEntry!]!
}

# A file, directory, or other tree entry.
interface TreeEntry {
  name: String!
  isDirectory: Boolean!
  repository: Repository!
  commits: [CommitInfo!]!
  lastCommit: CommitInfo!
}

type Directory implements TreeEntry {
  name: String!
  isDirectory: Boolean!
  repository: Repository!
  commits: [CommitInfo!]!
  lastCommit: CommitInfo!
  tree: Tree!
}

type HighlightedFile {
  aborted: Boolean!
  html: String!
}

type File implements TreeEntry {
  name: String!
  content: String!
  repository: Repository!
  binary: Boolean!
  isDirectory: Boolean!
  commit: Commit!
  highlight(disableTimeout: Boolean!): HighlightedFile!
  blame(startLine: Int!, endLine: Int!): [Hunk!]!
  commits: [CommitInfo!]!
  lastCommit: CommitInfo!
  dependencyReferences(
    Language: String!
    Line: Int!
    Character: Int!
  ): DependencyReferences!
  blameRaw(startLine: Int!, endLine: Int!): String!
}

type FileMatch {
  resource: String!
  lineMatches: [LineMatch!]!
  limitHit: Boolean!
}

type LineMatch {
  preview: String!
  lineNumber: Int!
  offsetAndLengths: [[Int!]!]!
  limitHit: Boolean!
}

type DependencyReferences {
  dependencyReferenceData: DependencyReferencesData!
  repoData: RepoDataMap!
}

type RepoDataMap {
  repos: [Repository!]!
  repoIds: [Int!]!
}

type DependencyReferencesData {
  references: [DependencyReference!]!
  location: DepLocation!
}

type DependencyReference {
  dependencyData: String!
  repoId: Int!
  hints: String!
}

type DepLocation {
  location: String!
  symbol: String!
}

type Hunk {
  startLine: Int!
  endLine: Int!
  startByte: Int!
  endByte: Int!
  rev: String!
  author: Signature
  message: String!
}

type Installation {
  login: String!
  githubId: Int!
  installId: Int!
  type: String!
  avatarURL: String!
}

type User implements Node, ConfigurationSubject {
  # The unique ID for the user.
  id: ID!
  auth0ID: String!
  sourcegraphID: Int
  email: String!
  displayName: String
  username: String
  avatarURL: String
  createdAt: String
  updatedAt: String
  # The latest settings for the user.
  latestSettings: Settings
  orgs: [Org!]!
  orgMemberships: [OrgMember!]!
  hasSourcegraphUser: Boolean!
  tags: [UserTag!]!
  activity: UserActivity!
}

type CompanyProfile {
  ip: String!
  domain: String!
  fuzzy: Boolean!
  company: CompanyInfo!
}

type CompanyInfo {
  id: String!
  name: String!
  legalName: String!
  domain: String!
  domainAliases: [String!]!
  url: String!
  site: SiteDetails!
  category: CompanyCategory!
  tags: [String!]!
  description: String!
  foundedYear: String!
  location: String!
  logo: String!
  tech: [String!]!
}

type SiteDetails {
  url: String!
  title: String!
  phoneNumbers: [String!]!
  emailAddresses: [String!]!
}

type CompanyCategory {
  sector: String!
  industryGroup: String!
  industry: String!
  subIndustry: String!
}

type Org implements Node, ConfigurationSubject {
  id: ID!
  orgID: Int!
  name: String!
  displayName: String
  slackWebhookURL: String
  members: [OrgMember!]!
  latestSettings: Settings
  repos: [OrgRepo!]!
  repo(canonicalRemoteID: String!): OrgRepo
  threads(
    repoCanonicalRemoteID: String
    branch: String
    file: String
    limit: Int
  ): ThreadConnection!
  tags: [OrgTag!]!
}

type OrgMember {
  id: Int!
  org: Org!
  user: User!
  createdAt: String!
  updatedAt: String!
}

type OrgInviteStatus {
  emailVerified: Boolean!
}

type OrgRepo {
  id: Int!
  org: Org!
  canonicalRemoteID: String!
  createdAt: String!
  updatedAt: String!
  threads(file: String, branch: String, limit: Int): ThreadConnection!
}

type ThreadConnection {
  nodes: [Thread!]!
  totalCount: Int!
}

# ConfigurationSubject is something that can have configuration.
interface ConfigurationSubject {
  id: ID!
  latestSettings: Settings
}

# The configurations for all of the relevant configuration subjects, plus the merged
# configuration.
type ConfigurationCascade {
  # The default settings, which are applied first and the lowest priority behind
  # all configuration subjects' settings.
  defaults: Configuration
  # The configurations for all of the subjects that are applied for the currently
  # authenticated user. For example, a user in 2 orgs would have the following
  # configuration subjects: org 1, org 2, and the user.
  subjects: [ConfigurationSubject!]!
  # The effective configuration, merged from all of the subjects.
  merged: Configuration!
}

# Settings is a version of a configuration settings file.
type Settings {
  id: Int!
  configuration: Configuration!
  # The subject that these settings are for.
  subject: ConfigurationSubject!
  author: User!
  createdAt: String!
  contents: String! @deprecated(reason: "use configuration.contents instead")
  highlighted: String!
    @deprecated(reason: "use configuration.highlighted instead")
}

# Configuration contains settings from (possibly) multiple settings files.
type Configuration {
  # The raw JSON contents, encoded as a string.
  contents: String!
  # The contents as highlighted HTML.
  highlighted: String!
  # Error and warning messages about the configuration.
  messages: [String!]!
}

type Thread {
  id: Int!
  repo: OrgRepo!
  file: String!
  branch: String
  # The commit ID of the repository at the time the thread was created.
  repoRevision: String!
  # The commit ID from Git blame, at the time the thread was created.
  #
  # The selection may be multiple lines, and the commit id is the
  # topologically most recent commit of the blame commit ids for the selected
  # lines.
  #
  # For example, if you have a selection of lines that have blame revisions
  # (a, c, e, f), and assuming a history like::
  #
  # 	a <- b <- c <- d <- e <- f <- g <- h <- HEAD
  #
  # Then lines_revision would be f, because all other blame revisions a, c, e
  # are reachable from f.
  #
  # Or in lay terms: "What is the oldest revision that I could checkout and
  # still see the exact lines of code that I selected?".
  linesRevision: String!
  title: String!
  startLine: Int!
  endLine: Int!
  startCharacter: Int!
  endCharacter: Int!
  rangeLength: Int!
  createdAt: String!
  archivedAt: String
  author: User!
  lines: ThreadLines
  comments: [Comment!]!
}

type Comment {
  id: Int!
  title: String!
  contents: String!
  createdAt: String!
  updatedAt: String!
  author: User!
}

type Package {
  lang: String!
  repo: Repository
  # The following fields are properties of build package configuration as returned by the workspace/xpackages LSP endpoint.
  id: String
  type: String
  name: String
  commit: String
  baseDir: String
  repoURL: String
  version: String
}

type Dependency {
  repo: Repository
  # The following fields are properties of build package configuration as returned by the workspace/xpackages LSP endpoint.
  name: String
  repoURL: String
  depth: Int
  vendor: Boolean
  package: String
  absolute: String
  type: String
  commit: String
  version: String
  id: String
  package: String
}

type UserTag {
  id: Int!
  name: String!
}

type OrgTag {
  id: Int!
  name: String!
}

type UserActivity {
  id: Int!
  searchQueries: Int!
  pageViews: Int!
  createdAt: String!
  updatedAt: String!
}

enum UserEvent {
  PAGEVIEW
  SEARCHQUERY
}
`
