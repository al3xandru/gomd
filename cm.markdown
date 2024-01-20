# Critic Markup

There are five additions from Critical Markdown. 

First is **addition** used {++to add new text++} which results in `<ins>`.

Second is **deletion** used {--to remove text--} which results in `<del>`.

Third is **substituion** used to {~~replace some text~>with new text~~} 
which results in a `<del>` followed by `<ins>`. 

Forth is **comment** used to {>> add comments that are visible <<} 
which is transformed into `<span class="critic comment">`.


Last, the **highlight** used {==to highlight some text==}{>>and add comments to it<<}
which becomes a `<mark>`, followed by the optional comment.


The rule for these elements is that space is significant:

*   {++ this addition has space at the beginning and end ++}
*   {-- this deletion has space at the beginning and end --}
*   {~~ this replacement has spaces surrounded removed text ~> and the new text ~~}

The following paragraphs are intended to show multi-line usage.

We start with an {++addition with
two newlines 
across three lines, but no new paragraphs++}.

We follow with a  {--deletion 
that 
spans
four lines--} to complete.

