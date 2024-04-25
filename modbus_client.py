from pymodbus.client import ModbusTcpClient

def run_modbus_client():
    # 连接到Modbus服务器
    client = ModbusTcpClient('localhost', port=5020)
    client.connect()

    # 写单个寄存器，地址为1，值为10
    client.write_register(1, 10)
    
    # 从同一寄存器读取值
    response = client.read_holding_registers(1, 1)
    if not response.isError():
        print("Register value:", response.registers[0])
    else:
        print("Error reading register")

    # 断开连接
    client.close()

if __name__ == "__main__":
    run_modbus_client()
