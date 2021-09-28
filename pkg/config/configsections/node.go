// Copyright (C) 2020-2021 Red Hat, Inc.
//
// This program is free software; you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation; either version 2 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along
// with this program; if not, write to the Free Software Foundation, Inc.,
// 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301 USA.

package configsections

// Node to store information about node. A node can be master and worker at the same time
type Node struct {
	// Name is the name of node
	Name string `yaml:"name" json:"name"`
	// Worker indicates if the node is defined as worker in OCP cluster
	Worker bool `yaml:"worker" json:"worker"`
	// Master indicates if the node is defined as master in OCP cluster
	Master bool `yaml:"master" json:"master"`
}
