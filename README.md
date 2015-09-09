/** **************************************************************************
 *  Copyright (c) 2015 Lobaro UG (haftungsbeschränkt)
 *  T.Rohde
 *
 *  ### Lobaro.com Simple UDP Test ###
 *
 *  Description:
 *  - This app sends out the current time as string to an given IPv6 address via UDP.
 *	- The interval for transmison of these udp packets is one second.
 * 	- The remote address and port must be set as commandline parameter.
 *	- Local port will be choosen randomly by the operating system.
 *	- Any received data will be simly shown on the command line as string.
 *
 *****************************************************************************/

usage: udptest.exe [IPv6_of_LobaroBox:Port]
Example: udptest.exe [fe80:0000:0000:0000:0211:7d00:0030:8e3f]:5684
