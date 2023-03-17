# wez-links

Create ExplicitHyperlinks from piped input for use with WezTerm hyperlink custom processing

## Brute force...

[WezTerm](https://wezfurlong.org/wezterm/) has a Hyperlink feature where you can 
define regexes and have the terminal examine the output and decorate identified 
matches with links.  Mostly launching external links, though one can turn those 
internal to perform a great deal of fun actions.  In this case I am using this 
method to simply "send_text" back to the command line as input for the next 
command I want to run.  Some example output of a `kubectl get secrets` command:

```plaintext 
NAME      TYPE                                  DATA   AGE
secret1   Opaque                                1      30d
secret2   kubernetes.io/service-account-token   3      30d
```

While I can certainly craft a regex to identify and apply an on-the-fly
hyperlink, the problem becomes that the entire match becomes the hyperlink.
What I would like to do, is use the first match `secret1`, `secret2` as data
for the hyperlink that is being built.  eg.

`secret1` and `secret2` above would both be hyperlinks that when clicked, simply
drop `secret1` and `secret2` back to the command line, using a "send_text"
function.

The complex part then, is to make `Opaque` and `kubernetes.io.service-account-token`
into hyperlinks that are also using `secret1` and `secret2` as part of the data
of the hyperlink itself.  eg.

`<bash:kubectl view-secret secret1 -a>` would be the hyperlink for `Opaque`

The `<bash:>` mechanism is my internal identifier for dropping things back to
the command line via the "send_text" mechanism. 

As the `short-term` brute force method, I'm simply looking at the first output 
line from the pipeline, matching it up to determine what `type` of input it is, 
then ripping it apart and applying the multiple hyperlinks to the output.

With only a few `types`, this will work for now.  It makes me wonder if there is 
a way to identify any of the upstream process names that are sending the input 
down the pipeline.  It would be cool if one could determine if the input was 
coming from `kubectl`, `gcloud`, or whatever, and then tailor the matching based
on that.  Likely one of those things I won't have time to get around to if this 
works well enough.

