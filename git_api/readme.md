### Github Repository Deleter Utility

This program deletes a specified repository in github. It takes three flag arguments at the command line:

-token: The text file containing the api token for github. Note it should have permission to delete repositories. 
Additional documentation here https://github.com/settings/tokens

-owner: The owner of the repository.

-repo: The name of the repository.

Example usage:

```
./git_repo_deleting_tool -token=token.txt -repo=test_repo -owner=JeremyFenwick
```
