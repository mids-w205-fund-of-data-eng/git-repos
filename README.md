# `git-repos`

git `repos` subcommand to help with grading github-classroom assignments and
managing org repos in bulk between semesters.


## Usage

Use this as

    git repos <subcommand>

where `subcommand` may be `list` or `reviews`.

with the following structure

    git repos list --org-name <org> <pattern>
    git repos reviews --org-name <org> --reviewer <reviewer> --state <review_state> <repo_name_pattern>



## Options

Note, The working org can either be specified from the `--org-name` option or
the `GITHUB_ORG` environment variable...

    GITHUB_ORG=mids-w205-martin-mims git repos list "*assignment-01*"

or even simply 

    GITHUB_ORG=mids-w205-martin-mims git repos list assignment-01


## Reviews

GitHub Classroom assignments each have a separate private repo per student assignment.
I.e., each student has a separate private repo per assignment.  Only the student
and the org owner (instructor) can read/write to each assignment repo.

In our classes:
- students submit assignments for grading by creating a pull request with the
  graders/instructors added as reviewers
- assignments are graded as reviews of these PRs, with the actual grade for the
  assignment recorded in the accompanying review comment

This tool is useful in tracking:

which students have accepted an assignment

    git repos list --org-name mids-w205-martin-mims assignment-01

which student assignments have been submitted for grading

    git repos reviews --org-name mids-w205-martin-mims --reviewer=mmm assignment-01

which student assignments have been graded and
the grades for each student assignment

    git repos reviews --org-name mids-w205-martin-mims --reviewer=mmm --state=APPROVED assignment-01

Additionally, this can be used by students to check grades

    git repos reviews --org-name mids-w205-martin-mims --reviewer=mmm --state=APPROVED <student-gh-id>

---

## Flush

Be careful here please!

Between semesters, we clean out instructor orgs.  In addition to the grading
subcommands listed above, this tool also provides a `flush` subcommand used to
delete student assignment repos by pattern:

    # Be careful here please!
    git repos flush --org-name <org> <pattern>

or really,

    git repos flush --org-name <org> --confirm <pattern>

Again, be careful!
