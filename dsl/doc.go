/*
Package dsl defines the DSL used by the slack bot to define commands.

Grammar summary

Below you can see the grammar needed in order for the bot commands to be
defined and functionality be implemented.

All capital letters tokens are tokens that have special meaning and have to
be present in order for the definition to take effect.

    separator = &&

    message = reaction | message | team_joined | channel_joined

    condition_term = content | user_id | channel_id
    condition_operator = is_true | is_false
    condition_value = regexp

    condition = condition_term condition_operator condition_value
    conditions = condition [separator condition]

    action = reply | react | dm | run_function

    message matches conditions then action action_value


Grammar description

| - logical OR separator

[ ... ] - allows for optional repetition of tokens

regexp - a special token which reflects a regular expression accepted by the go
standard library regexp package. for more information, visit:
https://godoc.org/regexp

separator - a special token value which allows different conditions to be
separated

conditions - a token which is formed by at least a single condition. if
multiple "condition" tokens are defined then they will be separated by using the
separator token. the separator will act as an and operation. this means that
in order for the terms component to evaluate to true it will need to return
true from every token

action - a token which defines what the bot will do in when the conditions
evaluate to true

reaction - a special token which represents a slack reaction

message - a special token which represents a standard slack message

team_joined - a special token which represents a team joined event

channel_joined - a special token which represents a channel joined event

message - a token which represents any of the possible values listed

reply - an action which means replying in the same slack channel as the event
that generated the trigger

react - respond with a slack reaction to a message

dm - respond with a slack message to the person which sent a slack message

run_function - run one of the builtin functions

action - defines one of the allowed functions that are allowed to be executed

matches - a special token that allows the user to express the intent of having
a message type matching to the expected conditions

then - a special token that allows the user to specify what happens when the
matching of the message is successful

action_value - value parameter for the action to use


Example

Below you can find examples of accepted rules by the above grammar:

    MESSAGE MATCHES CONTENT IS_TRUE newbie resources THEN REPLY here are some newbie resources...
    REACTION MATCHES CONTENT IS_TRUE gopherbaby THEN REACT gopherbaby

*/
package dsl
