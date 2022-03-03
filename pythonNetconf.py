import os
from logging.config import fileConfig
import concurrent.futures
import ipaddress
from jinja2 import Environment, FileSystemLoader
from ncclient import operations, manager
from lxml import etree
import logging
from datetime import datetime;
# [rest of code]
logging.basicConfig(format='%(asctime)s - %(message)s', filename='pythonNetconf.log', filemode='w', level=logging.INFO)
# logging.basicConfig(filename='example.log', filemode='w', level=logging.DEBUG)
file_loader = FileSystemLoader('templates')
env = Environment(loader=file_loader)
template = env.get_template('ncPayload.jinja')
def deployNetconf():
    logging.info("Started ")
    start = datetime.now()
    with concurrent.futures.ThreadPoolExecutor(max_workers=300) as executor:
        # executor.map(deployConfig, [str(ip) for ip in ipaddress.IPv4Network('192.0.0.0/24')])
        executor.map(deployConfig, ["10.2.31.2"])
        # executor.map(deployConfig, ['10.6.31.2', '10.6.31.9', '10.6.31.20', '10.6.31.3', '10.6.31.6', '10.6.31.11', '10.6.31.5', '10.6.31.13', '10.6.31.8', '10.6.31.7', '10.6.31.4', '10.6.31.50', '10.6.31.21', '10.6.31.39', '10.6.31.43', '10.6.31.44', '10.6.31.32', '10.6.31.46', '10.6.31.31', '10.6.31.34', '10.6.31.45', '10.6.31.35', '10.6.31.12', '10.6.31.17', '10.6.31.28', '10.6.31.23', '10.6.31.47', '10.6.31.22', '10.6.31.48', '10.6.31.14', '10.6.31.27', '10.6.31.10', '10.6.31.18', '10.6.31.42', '10.6.31.19', '10.6.31.52', '10.6.31.40', '10.6.31.16', '10.6.31.33', '10.6.31.49', '10.6.31.41', '10.6.31.29', '10.6.31.53', '10.6.31.25', '10.6.31.15', '10.6.31.37', '10.6.31.24', '10.6.31.26', '10.6.31.62', '10.6.31.56', '10.6.31.38', '10.6.31.57', '10.6.31.36', '10.6.31.61', '10.6.31.59', '10.6.31.55', '10.6.31.54', '10.6.31.51', '10.6.31.60', '10.6.30.5', '10.6.30.8', '10.6.30.9', '10.6.30.4', '10.6.30.27', '10.6.30.2', '10.6.30.49', '10.6.30.47', '10.6.30.87', '10.6.30.62', '10.6.30.66', '10.6.30.91', '10.6.30.56', '10.6.30.29', '10.6.30.58', '10.6.30.60', '10.6.30.79', '10.6.30.78', '10.6.30.14', '10.6.30.116', '10.6.30.85', '10.6.30.53', '10.6.30.42', '10.6.30.101', '10.6.30.99', '10.6.30.20', '10.6.30.17', '10.6.30.43', '10.6.30.35', '10.6.30.100', '10.6.30.21', '10.6.30.61', '10.6.30.37', '10.6.30.39', '10.6.30.34', '10.6.30.106', '10.6.30.41', '10.6.30.25', '10.6.30.52', '10.6.30.13', '10.6.30.59', '10.6.30.98', '10.6.30.80', '10.6.30.88', '10.6.30.18', '10.6.30.97', '10.6.30.48', '10.6.30.30', '10.6.30.94', '10.6.30.127', '10.6.30.81', '10.6.30.139', '10.6.30.112', '10.6.30.86', '10.6.30.125', '10.6.30.149', '10.6.30.119', '10.6.30.154', '10.6.30.122', '10.6.30.135', '10.6.30.144', '10.6.30.126', '10.6.30.117', '10.6.30.110', '10.6.30.152', '10.6.30.104', '10.6.30.179', '10.6.30.171', '10.6.30.153', '10.6.30.146', '10.6.30.121', '10.6.30.184', '10.6.30.185', '10.6.30.180', '10.6.30.183', '10.6.30.170', '10.6.30.151', '10.6.30.128', '10.6.30.172', '10.6.30.159', '10.6.30.168', '10.6.30.173', '10.6.30.169', '10.6.30.182', '10.6.30.134', '10.6.30.138', '10.6.30.113', '10.6.30.123', '10.6.31.71', '10.6.31.80', '10.6.31.67', '10.6.31.66', '10.6.31.89', '10.6.31.98', '10.6.31.96', '10.6.31.97', '10.6.31.81', '10.6.31.92', '10.6.31.95', '10.6.31.103', '10.6.31.110', '10.6.31.121', '10.6.31.118', '10.6.31.126', '10.6.31.122', '10.6.31.120', '10.6.31.113', '10.6.31.123', '10.6.31.112', '10.6.31.70', '10.6.31.90', '10.6.31.100', '10.6.31.124', '10.6.31.125', '10.6.31.128', '10.6.31.130', '10.6.31.132', '10.6.31.138', '10.6.31.133', '10.6.31.135', '10.6.31.63', '10.6.30.3', '10.6.30.10', '10.6.30.7', '10.6.30.6', '10.6.30.12', '10.6.30.72', '10.6.30.38', '10.6.30.68', '10.6.30.55', '10.6.30.64', '10.6.30.69', '10.6.30.32', '10.6.30.28', '10.6.30.45', '10.6.30.65', '10.6.30.50', '10.6.30.54', '10.6.30.63', '10.6.30.96', '10.6.30.16', '10.6.30.46', '10.6.30.82', '10.6.30.15', '10.6.30.44', '10.6.30.70', '10.6.30.31', '10.6.30.90', '10.6.30.51', '10.6.30.40', '10.6.30.73', '10.6.30.19', '10.6.30.24', '10.6.30.89', '10.6.30.95', '10.6.30.36', '10.6.30.76', '10.6.30.102', '10.6.30.57', '10.6.30.107', '10.6.30.77', '10.6.30.108', '10.6.30.124', '10.6.30.111', '10.6.30.33', '10.6.30.75', '10.6.30.84', '10.6.30.71', '10.6.30.26', '10.6.30.115', '10.6.30.133', '10.6.30.109', '10.6.30.67', '10.6.30.118', '10.6.30.83', '10.6.30.92', '10.6.30.158', '10.6.30.148', '10.6.30.120', '10.6.30.143', '10.6.30.105', '10.6.30.11', '10.6.30.136', '10.6.30.156', '10.6.30.147', '10.6.30.140', '10.6.30.142', '10.6.30.141', '10.6.30.175', '10.6.30.150', '10.6.30.132', '10.6.30.161', '10.6.30.162', '10.6.30.129', '10.6.30.131', '10.6.30.176', '10.6.30.177', '10.6.30.145', '10.6.30.165', '10.6.30.181', '10.6.30.160', '10.6.30.174', '10.6.30.163', '10.6.30.167', '10.6.30.114', '10.6.30.157', '10.6.31.84', '10.6.31.69', '10.6.31.79', '10.6.30.178', '10.6.30.137', '10.6.31.87', '10.6.31.88', '10.6.31.86', '10.6.31.83', '10.6.31.68', '10.6.31.77', '10.6.31.104', '10.6.31.94', '10.6.31.78', '10.6.31.101', '10.6.31.108', '10.6.31.107', '10.6.31.111', '10.6.31.116', '10.6.31.115', '10.6.31.72', '10.6.31.76', '10.6.31.64', '10.6.31.73', '10.6.31.93', '10.6.31.91', '10.6.31.114', '10.6.31.119', '10.6.31.74', '10.6.31.129', '10.6.31.139', '10.6.31.137', '10.6.31.134', '10.6.31.136', '10.6.31.75'])
        # executor.map(deployConfig, ['10.2.31.2', '10.2.31.9', '10.2.31.20', '10.2.31.3', '10.2.31.2', '10.2.31.11', '10.2.31.5', '10.2.31.13', '10.2.31.8', '10.2.31.7', '10.2.31.4', '10.2.31.50', '10.2.31.21', '10.2.31.39', '10.2.31.43', '10.2.31.44', '10.2.31.32', '10.2.31.46', '10.2.31.31', '10.2.31.34', '10.2.31.45', '10.2.31.35', '10.2.31.12', '10.2.31.17', '10.2.31.28', '10.2.31.23', '10.2.31.47', '10.2.31.22', '10.2.31.48', '10.2.31.14', '10.2.31.27', '10.2.31.10', '10.2.31.18', '10.2.31.42', '10.2.31.19', '10.2.31.52', '10.2.31.40', '10.2.31.16', '10.2.31.33', '10.2.31.49', '10.2.31.41', '10.2.31.29', '10.2.31.53', '10.2.31.25', '10.2.31.15', '10.2.31.37', '10.2.31.24', '10.2.31.26', '10.2.31.22', '10.2.31.56', '10.2.31.38', '10.2.31.57', '10.2.31.36', '10.2.31.21', '10.2.31.59', '10.2.31.55', '10.2.31.54', '10.2.31.51', '10.2.31.20', '10.2.30.5', '10.2.30.8', '10.2.30.9', '10.2.30.4', '10.2.30.27', '10.2.30.2', '10.2.30.49', '10.2.30.47', '10.2.30.87', '10.2.30.22', '10.2.30.26', '10.2.30.91', '10.2.30.56', '10.2.30.29', '10.2.30.58', '10.2.30.20', '10.2.30.79', '10.2.30.78', '10.2.30.14', '10.2.30.116', '10.2.30.85', '10.2.30.53', '10.2.30.42', '10.2.30.101', '10.2.30.99', '10.2.30.20', '10.2.30.17', '10.2.30.43', '10.2.30.35', '10.2.30.100', '10.2.30.21', '10.2.30.21', '10.2.30.37', '10.2.30.39', '10.2.30.34', '10.2.30.106', '10.2.30.41', '10.2.30.25', '10.2.30.52', '10.2.30.13', '10.2.30.59', '10.2.30.98', '10.2.30.80', '10.2.30.88', '10.2.30.18', '10.2.30.97', '10.2.30.48', '10.2.30.30', '10.2.30.94', '10.2.30.127', '10.2.30.81', '10.2.30.139', '10.2.30.112', '10.2.30.86', '10.2.30.125', '10.2.30.149', '10.2.30.119', '10.2.30.154', '10.2.30.122', '10.2.30.135', '10.2.30.144', '10.2.30.126', '10.2.30.117', '10.2.30.110', '10.2.30.152', '10.2.30.104', '10.2.30.179', '10.2.30.171', '10.2.30.153', '10.2.30.146', '10.2.30.121', '10.2.30.184', '10.2.30.185', '10.2.30.180', '10.2.30.183', '10.2.30.170', '10.2.30.151', '10.2.30.128', '10.2.30.172', '10.2.30.159', '10.2.30.168', '10.2.30.173', '10.2.30.169', '10.2.30.182', '10.2.30.134', '10.2.30.138', '10.2.30.113', '10.2.30.123', '10.2.31.71', '10.2.31.80', '10.2.31.27', '10.2.31.26', '10.2.31.89', '10.2.31.98', '10.2.31.96', '10.2.31.97', '10.2.31.81', '10.2.31.92', '10.2.31.95', '10.2.31.103', '10.2.31.110', '10.2.31.121', '10.2.31.118', '10.2.31.126', '10.2.31.122', '10.2.31.120', '10.2.31.113', '10.2.31.123', '10.2.31.112', '10.2.31.70', '10.2.31.90', '10.2.31.100', '10.2.31.124', '10.2.31.125', '10.2.31.128', '10.2.31.130', '10.2.31.132', '10.2.31.138', '10.2.31.133', '10.2.31.135', '10.2.31.23', '10.2.30.3', '10.2.30.10', '10.2.30.7', '10.2.30.2', '10.2.30.12', '10.2.30.72', '10.2.30.38', '10.2.30.28', '10.2.30.55', '10.2.30.24', '10.2.30.29', '10.2.30.32', '10.2.30.28', '10.2.30.45', '10.2.30.25', '10.2.30.50', '10.2.30.54', '10.2.30.23', '10.2.30.96', '10.2.30.16', '10.2.30.46', '10.2.30.82', '10.2.30.15', '10.2.30.44', '10.2.30.70', '10.2.30.31', '10.2.30.90', '10.2.30.51', '10.2.30.40', '10.2.30.73', '10.2.30.19', '10.2.30.24', '10.2.30.89', '10.2.30.95', '10.2.30.36', '10.2.30.76', '10.2.30.102', '10.2.30.57', '10.2.30.107', '10.2.30.77', '10.2.30.108', '10.2.30.124', '10.2.30.111', '10.2.30.33', '10.2.30.75', '10.2.30.84', '10.2.30.71', '10.2.30.26', '10.2.30.115', '10.2.30.133', '10.2.30.109', '10.2.30.27', '10.2.30.118', '10.2.30.83', '10.2.30.92', '10.2.30.158', '10.2.30.148', '10.2.30.120', '10.2.30.143', '10.2.30.105', '10.2.30.11', '10.2.30.136', '10.2.30.156', '10.2.30.147', '10.2.30.140', '10.2.30.142', '10.2.30.141', '10.2.30.175', '10.2.30.150', '10.2.30.132', '10.2.30.161', '10.2.30.162', '10.2.30.129', '10.2.30.131', '10.2.30.176', '10.2.30.177', '10.2.30.145', '10.2.30.165', '10.2.30.181', '10.2.30.160', '10.2.30.174', '10.2.30.163', '10.2.30.167', '10.2.30.114', '10.2.30.157', '10.2.31.84', '10.2.31.29', '10.2.31.79', '10.2.30.178', '10.2.30.137', '10.2.31.87', '10.2.31.88', '10.2.31.86', '10.2.31.83', '10.2.31.28', '10.2.31.77', '10.2.31.104', '10.2.31.94', '10.2.31.78', '10.2.31.101', '10.2.31.108', '10.2.31.107', '10.2.31.111', '10.2.31.116', '10.2.31.115', '10.2.31.72', '10.2.31.76', '10.2.31.24', '10.2.31.73', '10.2.31.93', '10.2.31.91', '10.2.31.114', '10.2.31.119', '10.2.31.74', '10.2.31.129', '10.2.31.139', '10.2.31.137', '10.2.31.134', '10.2.31.136', '10.2.31.75'])

    logging.info("Ended ")
    end=datetime.now()
    print("time taken to complete " + str(end - start))

def deployConfig(ipAddress):
    requestBody = template.render(ipAddress=ipAddress)
    logging.info(""+ipAddress)
    # logging.info(requestBody)
    m = None




    try:
        m = manager.connect(host=ipAddress, port=830,
                            username='admin', password='admin',
                            hostkey_verify=False, timeout = 10)
        
        logging.info("after connected "+ipAddress)

        rpc_elem = etree.fromstring(requestBody.encode("utf-8"))
        reply = m.rpc(rpc_elem)
        rpc_elem = etree.fromstring("<commit/>".encode("utf-8"))
        reply=m.rpc(rpc_elem)
        # print("result  : ", reply)
        m.close_session()
    except Exception as err:
        print("Error ", str(err))
        m.close_session()

if __name__ == "__main__":
    deployNetconf()