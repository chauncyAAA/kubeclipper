import yaml
import time
import subprocess
from yaml.loader import SafeLoader

start_shell = "gabbi-run -v all 172.20.139.223 -- {}"


result_data = {}
totoal_api = 89
current_api = 34
total_number = 0
pass_number = 0
fail_number = 0


with open("list.yaml") as f:
    list_data = yaml.load(f, Loader=SafeLoader)


def runcmd(command, num, output=False):
    if output:
        print(command)
    command_list = command.split()
    p =subprocess.Popen(command_list, shell=False, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
    _, stderr = p.communicate()
    time.sleep(1)
    result = stderr.decode("utf-8")
    return result.split("\n")[:num+1]


def create_test_cluster(retry):
    if retry <= 0:
        return 0
    result = runcmd(start_shell.format("test_cluster/check_test_cluster.yaml"),1)
    if result[1][4]!="✓":
        print(result)
        runcmd(start_shell.format("test_cluster/delete_test_cluster.yaml"),1)
        result = runcmd(start_shell.format("test_cluster/create_test_cluster.yaml"),1)
        if result[1][4]!="✓":
            print(result)
            create_test_cluster(retry-1)
    return 1


for list_project in list_data["apitest"]:
    if create_test_cluster(3) != 1:
        continue
    project = list(list_project.keys())[0]
    result_data[project] = {}
    for list_subproject in list_project[project]:
        subproject = list(list_subproject.keys())[0]
        result_data[project][subproject] = {}
        file_name = project + '/' + list_subproject[subproject]["file_name"]
        number = list_subproject[subproject]["number"]
        total_number += number
        result = runcmd(start_shell.format(file_name), number, True)
        for res in result[1:number+1]:
            task = res[20+len(subproject):]
            if res[4] == "✓":
                pass_number += 1
                result_data[project][subproject][task] = 'pass'
            else:
                fail_number += 1
                print(res)
                result_data[project][subproject][task] = 'fail'


with open("report","w+") as f:
    f.write("total number: {}\n".format(total_number))
    f.write("pass number: {}\n".format(pass_number))
    f.write("fail number: {}\n".format(fail_number))
    f.write("pass coverage: {}%\n".format(round((pass_number*1.0/total_number), 2)*100))
    f.write("api coverage: {}%\n".format(round((current_api*1.0/totoal_api), 2)*100))
    f.write("---------------------------------------------------------------\n")
    f.write("Details:\n")
    for project in list(result_data.keys()):
        f.write(" {}:\n".format(project))
        for subproject in result_data[project].keys():
            f.write("\t{}:\n".format(subproject))
            for task in result_data[project][subproject].keys():
                f.write("\t\t{}: {}\n".format(task, result_data[project][subproject][task]))


if pass_number==total_number:
    with open("result","w+") as f:
        f.write("api-test successfully")