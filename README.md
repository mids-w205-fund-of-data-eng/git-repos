# `git-repos`

Manage org repos in bulk.


## Usage

Use this as

    git repos <subcommand>

where `subcommand` may be

    list | flush

with the following structure

    git repos list --org-name <org> <pattern>
    git repos flush --org-name <org> --confirm <pattern>


Note, The working org can either be specified from the `--org-name` option or
the `GITHUB_ORG` environment variable...

    GITHUB_ORG=mids-w205-martin-mims git repos list "*assignment-01*"

or even simply 

    GITHUB_ORG=mids-w205-martin-mims git repos list assignment-01

and then

    GITHUB_ORG=mids-w205-martin-mims git repos flush assignment-01
    GITHUB_ORG=mids-w205-martin-mims git repos flush --confirm assignment-01


---

## Reviews

    git repos reviews --org-name <org> --reviewer <reviewer> --state <review_state> <repo_name_pattern>

for example,

    git repos reviews --org-name mids-w205-martin-mims --reviewer=mmm --state=APPROVED assignment-01
