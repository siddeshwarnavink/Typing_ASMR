import os
import shutil
import subprocess
import sys

def build_go_project():
    print("Building the Go project...")
    subprocess.check_call(["go", "build", "-o", "typing-asmr"])

def create_installation_directory():
    install_dir = "/usr/local/typing-asmr" if os.name != 'nt' else os.path.join(os.environ['ProgramFiles'], 'typing-asmr')
    print(f"Creating installation directory at {install_dir}...")
    os.makedirs(install_dir, exist_ok=True)
    return install_dir

def copy_files(install_dir):
    print("Copying files to the installation directory...")
    shutil.copy("typing-asmr", install_dir)
    shutil.copytree("sounds", os.path.join(install_dir, "sounds"), dirs_exist_ok=True)

def update_path(install_dir):
    if os.name == 'nt':
        print(f"Adding {install_dir} to the PATH...")
        os.system(f'setx PATH "%PATH%;{install_dir}"')
    else:
        print(f"Adding {install_dir} to the PATH...")
        bashrc_path = os.path.expanduser("~/.bashrc")
        with open(bashrc_path, "a") as bashrc:
            bashrc.write(f'\nexport PATH=$PATH:{install_dir}')
        print(f"Run 'source {bashrc_path}' or open a new terminal to update PATH.")

def main():
    if os.name != 'nt' and os.geteuid() != 0:
        print("This script needs to be run with elevated privileges. Please use 'sudo'.")
        sys.exit(1)
    try:
        build_go_project()
        install_dir = create_installation_directory()
        copy_files(install_dir)
        update_path(install_dir)
        print("Installation complete!")
        print("You can now run 'typing-asmr' from anywhere.")
    except PermissionError:
        print("Permission denied. Please run this script with elevated privileges (e.g., using 'sudo').")
    except Exception as e:
        print(f"An error occurred: {e}")

if __name__ == "__main__":
    main()
