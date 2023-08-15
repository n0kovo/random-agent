# random-agent
 Simply output a random user-agent (with match and filter arguments).
 
 Pretty handy for tools without a `--random-agent` feature:
 
- `curl -A "$(random-agent)" https://example.com`
- `ffuf -w wordlist.txt https://example.com/FUZZ -H "User-Agent: $(random-agent -m chrome)"`
- `wfuzz -w common.txt --hc 404 -H "$(random-agent -m chrome -f android)" http://testphp.vulnweb.com/FUZZ`
- and so on :)

