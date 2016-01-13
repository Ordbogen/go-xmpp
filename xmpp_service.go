// XEP-0133: Service Administration
package xmpp

import (
	"fmt"
)

// Send announcement to all online users.
// Announcement will be sent only to users who currently
// have a "session" with the service.
// (http://xmpp.org/extensions/xep-0133.html#announce)
func (c *Client) SendAnnouncementToOnlineUsers(toServer, body string) {
	fmt.Fprintf(c.conn, "<iq to='%s' type='set' xml:lang='en'>"+
		"<command xmlns='http://jabber.org/protocol/commands' node='http://jabber.org/protocol/admin#announce'>"+
		"<x xmlns='jabber:x:data' type='submit'>"+
		"<field type='hidden' var='FORM_TYPE'>"+
		"<value>http://jabber.org/protocol/admin</value>"+
		"</field>"+
		"<field var='announcement'>"+
		"<value>%s</value>"+
		"</field>"+
		"<field var='body'>"+
		"<value>%s</value>"+
		"</field>"+
		"</x>"+
		"</command>"+
		"</iq>",
		xmlEscape(toServer), xmlEscape(body), body)
}
