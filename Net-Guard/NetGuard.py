import subprocess
import argparse
import time
import pygame
import threading
from termcolor import colored

pygame.mixer.init(frequency=44100, size=-16, channels=2, buffer=512)
pygame.mixer.set_num_channels(32)

soundPath = "/home/lain/Music/factorio-attack-alert.wav"
VPNPath   = "/home/lain/vpnbook-openvpn-de20/vpnbook-de20-tcp443.ovpn"

myIP = ""

sound1 = pygame.mixer.Sound(soundPath)
sound2 = pygame.mixer.Sound(soundPath)
sound3 = pygame.mixer.Sound(soundPath)
sound4 = pygame.mixer.Sound(soundPath)

sound1.set_volume(1.0)
sound2.set_volume(1.0)
sound3.set_volume(1.0)
sound4.set_volume(1.0)

TimeToCheckNet   = 10
TimeToCheckIPv6  = 180
TimeToCheckProxy = 60
TimeToCheckVPN   = 30

def checkNetwork():
	while True:
		try:
			output =  subprocess.check_output(['ip', 'a'], text=True)
			output1 = subprocess.check_output(['ping', '-c', '1', 'google.com'], text=True)
			
			if "inet 192.168" in output and "ms" in output1:
				sound1.stop()
				print(colored('\n[*]','blue'), 'The network is working correctly')
			else:
				sound1.play(-1)
				print(colored('\n[!]','red'), 'The network is lost!')
				print(colored('\n[*]','blue'), 'Starting the restart process...')
				
				subprocess.run(['sudo', 'systemctl', 'restart', 'NetworkManager'], text=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
				time.sleep(3)
				
				output = subprocess.check_output(['ip', 'a'], text=True)
				if "inet 192.168" in output:
					sound1.stop()
					print(colored('\n[+]','yellow'), 'The network has been successfully resumed!')
				else:
					print(colored('\n[-]','red'), 'Network restart error!')

		
		except:
			sound1.play(-1)
			print(colored('\n[!]','red'), 'The network is lost!')
			print(colored('\n[*]','blue'), 'Starting the restart process...')
			
			subprocess.run(['sudo', 'systemctl', 'restart', 'NetworkManager'], text=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
			time.sleep(3)
			
			output = subprocess.check_output(['ip', 'a'], text=True)
			if "inet 192.168" in output:
				sound1.stop()
				print(colored('\n[+]','yellow'), 'The network has been successfully resumed!')
			else:
				print(colored('\n[-]','red'), 'Network restart error!')
			
		time.sleep(TimeToCheckNet)

def ipv6Verification():
	while True:
		try:
			outputIPv6 = subprocess.run(['curl', '-6', 'https://icanhazip.com'], text=True, capture_output=True)
			time.sleep(1)
			if "Could not resolve host: icanhazip.com" in outputIPv6.stderr or "Failed to connect to icanhazip.com" in outputIPv6.stderr:
				print(colored('\n[*]','blue'), 'IPv6 is hidden')
			else:
				sound2.play(-1)
				print(colored('\n[!]','red'), 'IPv6 is visible!')
				print(colored('\n[*]','green'), 'Starting the IPv6 hiding process...')
				subprocess.run(['sudo', 'sysctl', '-p'], text=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
				outputIPv6 = subprocess.run(['curl', '-6', 'https://icanhazip.com'], text=True, capture_output=True)
				time.sleep(1)
				
				if "curl: (6) Could not resolve host: icanhazip.com" in outputIPv6.stderr:
					sound2.stop()
					print(colored('\n[+]','yellow'), 'IPv6 successfully hidden!')

		except subprocess.CalledProcessError:
			sound2.play(-1)
			print(colored('\n[-]','red'), 'Unknown error')
			
		time.sleep(TimeToCheckIPv6)

def CurlIP():
	while True:
		try:
			myIP = subprocess.run(['curl', 'https://icanhazip.com'], text=True, capture_output=True, timeout = 5)
			hiddenIP = subprocess.run(['proxychains' ,'curl', 'https://icanhazip.com'], text=True, capture_output=True, timeout = 5)
			
			if myIP.stderr not in hiddenIP.stderr:
				print(colored('\n[*]','blue'), 'Proxychains is working correctly')
			else:
				sound3.play(-1)
				print(colored('\n[!]','red'), 'Proxychains or curl timed out!')
				print(colored('\n[*]','blue'), 'Starting the restart process...')
				
				subprocess.run(['sudo', 'systemctl', 'restart', 'NetworkManager'], text=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
				time.sleep(3)
				
				myIP = subprocess.run(['curl', 'https://icanhazip.com'], text=True, capture_output=True, timeout = 5)
				hiddenIP = subprocess.run(['proxychains' ,'curl', 'https://icanhazip.com'], text=True, capture_output=True, timeout = 5)
				
				if myIP.stderr not in hiddenIP.stderr:
					sound3.stop()
					print(colored('\n[*]','blue'), 'The network has been successfully rebooted and the proxy is working.')
				else:
					print(colored('\n[!]','red'), 'Restart error!')
				
		except subprocess.TimeoutExpired:
			sound3.play(-1)
			print(colored('\n[!]','red'), 'Proxychains or curl timed out!')
			print(colored('\n[*]','blue'), 'Starting the restart process...')
			
			subprocess.run(['sudo', 'systemctl', 'restart', 'NetworkManager'], text=True, stdout=subprocess.DEVNULL, stderr=subprocess.DEVNULL)
			time.sleep(3)
			
			myIP = subprocess.run(['curl', 'https://icanhazip.com'], text=True, capture_output=True, timeout = 5)
			hiddenIP = subprocess.run(['proxychains' ,'curl', 'https://icanhazip.com'], text=True, capture_output=True, timeout = 5)
			
			if myIP.stderr not in hiddenIP.stderr:
				sound3.stop()
				print(colored('\n[*]','blue'), 'The network has been successfully rebooted and the proxy is working.')
			else:
				print(colored('\n[!]','red'), 'Restart error!')
			
				
		except subprocess.CalledProcessError:
			sound3.play(-1)
			print(colored('\n[-]','red'), 'Unknown error')
			
		time.sleep(TimeToCheckProxy)

def OpenVPN():
	global myIP 
	myIP = subprocess.run(['curl', '-4', 'https://icanhazip.com'], text=True, capture_output=True, timeout = 5)
	myIP = myIP.stdout.strip()  # оставить только IP без лишнего
	
	print(colored('\n[*]','blue'), 'Your current IPv4:', myIP)
	
	print(colored('\n[*]','green'), 'Opening the VPN...')
	
	output = subprocess.run(['sudo', 'openvpn', '--config', VPNPath, '--daemon'], text=True, capture_output=True)
	time.sleep(15)
	
	newIP = subprocess.run(['curl', '-4', 'https://icanhazip.com'], text=True, capture_output=True, timeout = 5)
	newIP = newIP.stdout.strip()
	
	if myIP not in newIP:
		print(colored('\n[+]','yellow'), 'VPN successfully connected')
		print(colored('\n[*]','blue'), 'Your new fake IPv4:', newIP)
	else:
		print(colored('\n[-]','red'), 'Error')

def CheckVPN():
	while True:
		try:
			global myIP
			
			if myIP == '':
				print(colored('\n[!]','cyan'), 'Before checking if the VPN is working correctly, enter your real external IP')
				
				myIP = input("\n> ")
			else:
				fakeIP = subprocess.run(['curl', '-4', 'https://icanhazip.com'], text=True, capture_output=True, timeout = 15)
				fakeIP = fakeIP.stdout.strip()
				
				if myIP not in fakeIP:
					sound4.stop()
					print(colored('\n[*]','blue'), 'VPN is working correctly')
				else:
					sound4.play(-1)
					print(colored('\n[!]','red'), 	'VPN not working!')
					print(colored('\n[*]','green'), 'Trying to reboot openvpn...')
					
					subprocess.run(['sudo', 'systemctl', 'restart', 'openvpn'], text=True, capture_output=True, timeout = 5)
					time.sleep(1.5)
				
		except subprocess.TimeoutExpired:
				sound4.play(-1)
				print(colored('\n[!]','red'), 'Curl timed out!')
				
		except subprocess.CalledProcessError:
				sound4.play(-1)
				print(colored('\n[-]','red'), 'Unknown error')
				
		time.sleep(TimeToCheckVPN)

def main():
	
	parser = argparse.ArgumentParser(description="a project for monitoring the correct operation of programs.")
	
	parser.add_argument("-a", 			"--all", 				help="run everything", 									action="store_true")
	parser.add_argument("-cN", 			"--check-net", 			help="check if the network is working correctly", 		action="store_true")
	parser.add_argument("-c6", 			"--check-ipv6", 		help="check if IPv6 is masked", 						action="store_true")
	parser.add_argument("-cP", 			"--check-proxy",		help="check if proxychains is masking your IP", 		action="store_true")
	parser.add_argument("-oVPN", 		"--open-vpn", 			help="connect to vpn", 									action="store_true")
	parser.add_argument("-cVPN", 		"--check-vpn", 			help="check if the VPN is working correctly", 			action="store_true")
	
	args = parser.parse_args()
	
	t1 = threading.Thread(target=checkNetwork)
	t2 = threading.Thread(target=ipv6Verification)
	t3 = threading.Thread(target=CurlIP)
	t4 = threading.Thread(target=CheckVPN)
	
	if not any(vars(args).values()):
	    parser.print_help()
	    exit()
	
	if args.check_net:
		t1.start()
		
	if args.check_ipv6:
		t2.start()
		
	if args.check_proxy:
		t3.start()
	
	if args.check_vpn:
		t4.start()
	
	if args.open_vpn:
		OpenVPN()
	
	if args.all:
		OpenVPN()
		t1.start()
		t2.start()
		t3.start()
		t4.start()
		
if __name__ == "__main__":
	main()
