---
phase: 01-backend-cleanup
plan: 02
type: execute
wave: 1
depends_on: []
files_modified:
  - server/plugin/register.go
autonomous: true
requirements: [CLEAN-07, CLEAN-08]
must_haves:
  truths:
    - "server/plugin/announcement/ directory does not exist"
    - "server/plugin/email/ directory does not exist"
    - "server/plugin/register.go does not reference announcement plugin"
  artifacts:
    - path: "server/plugin/announcement/"
      ensures: "Announcement plugin removed"
    - path: "server/plugin/email/"
      ensures: "Email plugin removed"
    - path: "server/plugin/register.go"
      ensures: "Plugin registration cleaned"
  key_links:
    - from: "server/plugin/register.go"
      to: "announcement plugin"
      pattern: "announcement"
---

## Plan 01-02: Delete Plugin Directories

### Objective
Delete server/plugin/announcement/ and server/plugin/email/ directories and clean up plugin registration.

### Context

@server/plugin/register.go (current state):
```go
package plugin

import (
    _ "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement"
)
```

### Tasks

#### Task 1: Remove plugin imports from register.go

<read_first>
- server/plugin/register.go
</read_first>

<action>
Update server/plugin/register.go to remove the announcement plugin import.
After deletion, the file should be empty or contain only package declaration with no imports.

Current content:
```go
package plugin

import (
    _ "github.com/flipped-aurora/gin-vue-admin/server/plugin/announcement"
)
```

Should become:
```go
package plugin
```

Or be deleted entirely if empty.
</action>

<acceptance_criteria>
- server/plugin/register.go no longer contains any plugin imports
- No compilation errors
</acceptance_criteria>

#### Task 2: Delete server/plugin/announcement/ directory

<read_first>
- Verify announcement plugin structure exists
</read_first>

<action>
Delete the entire server/plugin/announcement/ directory and all its contents:
- server/plugin/announcement/api/ (entire directory)
- server/plugin/announcement/config/ (entire directory)
- server/plugin/announcement/initialize/ (entire directory)
- server/plugin/announcement/model/ (entire directory)
- server/plugin/announcement/router/ (entire directory)
- server/plugin/announcement/service/ (entire directory)
- server/plugin/announcement/plugin/ (entire directory)
- server/plugin/announcement/plugin.go
</action>

<acceptance_criteria>
- server/plugin/announcement/ directory does not exist
- ls server/plugin/ shows only: email/ plugin-tool/ register.go
</acceptance_criteria>

#### Task 3: Delete server/plugin/email/ directory

<read_first>
- Verify email plugin structure exists
</read_first>

<action>
Delete the entire server/plugin/email/ directory and all its contents:
- server/plugin/email/api/ (entire directory)
- server/plugin/email/config/ (entire directory)
- server/plugin/email/initialize/ (entire directory)
- server/plugin/email/model/ (entire directory)
- server/plugin/email/router/ (entire directory)
- server/plugin/email/service/ (entire directory)
- server/plugin/email/plugin/ (entire directory)
- server/plugin/email/plugin.go
</action>

<acceptance_criteria>
- server/plugin/email/ directory does not exist
- ls server/plugin/ shows only: plugin-tool/ register.go
</acceptance_criteria>

### Verification

Run these commands to verify deletion:
```bash
# Verify plugin directories removed
ls server/plugin/   # Should show: plugin-tool/ register.go (and no email/ or announcement/)

# Verify Go files compile
cd server && go build ./...
```

### Success Criteria

1. server/plugin/announcement/ does not exist
2. server/plugin/email/ does not exist
3. server/plugin/register.go is empty or contains no plugin imports
4. Go code compiles without errors

### Output

After completion, create `.planning/phases/01-backend-cleanup/01-02-delete-plugin-dirs-SUMMARY.md`
