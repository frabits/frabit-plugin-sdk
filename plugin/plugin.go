/* (c) 2022 Frabit Project maintained and limited by Blylei < blylei.info@gmail.com >
GNU General Public License v3.0+ (see COPYING or https://www.gnu.org/licenses/gpl-3.0.txt)

This file is part of Frabit

*/

package plugin

type Plugin struct {
	Name        string
	Description string
	License     string
	Vendor      string
}

// Validate validate this plugin is a real plugin used with main application before install is
func (p *Plugin) Validate() bool {

	return true
}
