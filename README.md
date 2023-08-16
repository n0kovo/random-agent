# random-agent

### What?
 This simply outputs a random user-agent (with match and filter arguments).

### Why?
 Many web security tools have a `--random-agent` feature.<br>
 However, some don't, and I find that annoying.<br>
 This solves that problem ✨

### How?
- ```shell
  curl -A "`random-agent`" https://example.com
  ```

- ```shell
  ffuf -w wordlist.txt https://example.com/FUZZ -H "User-Agent: $(random-agent -m chrome)"
  ```

- ```shell
  wfuzz -w common.txt --hc 404 -H "$(random-agent -m chrome -f android)" http://testphp.vulnweb.com/FUZZ
  ```
and so on :)

- `-m string` / `--match string` only returns user agents matching `string`
- `-f string` / `--filter string` excludes user agents matching `string`
|Both arguments can be passed several times.

### Cool!
Right?? Run this and you're good to Go:
```shell
go install github.com/n0kovo/random-agent@latest
```

> [!NOTE]
> - The user-agent.txt file is embedded in the binary on build, so it's fully portable. If you want to edit or provide your own user-agent file, just clone the repo, edit `user-agents.txt` and run `go build .`.
> - Hack the planet
