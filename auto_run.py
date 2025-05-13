import subprocess
import os
import sys
import time

def main():
    project_root = os.path.dirname(os.path.abspath(__file__))
    print(f"[INFO] Project root: {project_root}")

    # --- 1. Start Docker Infrastructure ---
    print("\n[INFO] Starting Docker infrastructure...")
    infra_dir = os.path.join(project_root, "infrastructure")
    docker_compose_file = os.path.join(infra_dir, "docker-compose.yml")

    if not os.path.exists(docker_compose_file):
        print(f"[ERROR] docker-compose.yml not found at {docker_compose_file}")
        sys.exit(1)

    try:
        # Try 'docker compose' (newer syntax)
        subprocess.run(["docker", "compose", "up", "-d"], cwd=infra_dir, check=True, capture_output=True, text=True)
        print("[INFO] Docker infrastructure (using 'docker compose') started in detached mode.")
    except (FileNotFoundError, subprocess.CalledProcessError) as e_compose:
        print(f"[INFO] 'docker compose' command failed or not found ({e_compose}), trying 'docker-compose'...")
        try:
            # Fallback to 'docker-compose' (older syntax)
            subprocess.run(["docker-compose", "up", "-d"], cwd=infra_dir, check=True, capture_output=True, text=True)
            print("[INFO] Docker infrastructure (using 'docker-compose') started in detached mode.")
        except FileNotFoundError:
            print("[ERROR] Docker Compose command ('docker compose' or 'docker-compose') not found. Please ensure Docker is installed and in your PATH.")
            sys.exit(1)
        except subprocess.CalledProcessError as e_compose_old:
            print(f"[ERROR] Failed to start Docker infrastructure with 'docker-compose': {e_compose_old}")
            print(f"[ERROR] STDOUT: {e_compose_old.stdout}")
            print(f"[ERROR] STDERR: {e_compose_old.stderr}")
            sys.exit(1)
    except Exception as e: # Catch any other unexpected errors
        print(f"[ERROR] An unexpected error occurred with Docker Compose: {e}")
        sys.exit(1)

    # Brief pause to let Docker services initialize
    print("[INFO] Waiting 5 seconds for Docker services to initialize...")
    time.sleep(5)

    # --- 2. Run Go Project ---
    print("\n[INFO] Starting Go application...")
    cmd_dir = os.path.join(project_root, "cmd")
    go_main_file = os.path.join(cmd_dir, "main.go")
    log_file_path = os.path.join(project_root, "simple_ewallet_server.log")

    if not os.path.exists(go_main_file):
        print(f"[ERROR] Go main file not found at {go_main_file}")
        sys.exit(1)

    # Command to run the Go application.
    # If you prefer 'air' for live reloading, you can change this to:
    # go_run_cmd = ["air"]
    # Ensure 'air' is installed and configured in your 'cmd' directory if you use it.
    go_run_cmd = ["go", "run", "main.go"]
    
    print(f"[INFO] Go application output will be logged to: {log_file_path}")
    print(f"[INFO] To stop the Go application, press Ctrl+C in this terminal.")
    
    go_process = None
    try:
        # Open the log file in write mode to overwrite previous logs on each run.
        # Use 'a' for append mode if you prefer to keep old logs.
        with open(log_file_path, 'w') as lf:
            go_process = subprocess.Popen(
                go_run_cmd, 
                cwd=cmd_dir, 
                stdout=lf, 
                stderr=subprocess.STDOUT, # Redirect stderr to stdout, so both go to the log file
                text=True # Decode output as text
            )
        
        print(f"[INFO] Go application started (PID: {go_process.pid}). Check {log_file_path} for logs.")
        
        # Wait for the Go process to complete. This makes the script block here.
        # If the Go app exits or is terminated (e.g., by Ctrl+C), wait() will return.
        go_process.wait()

    except FileNotFoundError:
        print(f"[ERROR] Command '{go_run_cmd[0]}' not found. Make sure Go (or 'air' if you configured it) is installed and in your PATH.")
        if go_process and go_process.poll() is None: # If process started but command within failed (unlikely for 'go run' itself)
            go_process.kill()
        sys.exit(1)
    except KeyboardInterrupt:
        print("\n[INFO] KeyboardInterrupt received. Attempting to stop Go application...")
        if go_process and go_process.poll() is None: # Check if process exists and is running
            go_process.terminate() # Send SIGTERM for graceful shutdown
            try:
                go_process.wait(timeout=10) # Wait up to 10 seconds
                print("[INFO] Go application terminated gracefully.")
            except subprocess.TimeoutExpired:
                print("[INFO] Go application did not terminate gracefully after 10s, killing...")
                go_process.kill() # Force kill
                print("[INFO] Go application killed.")
        else:
            print("[INFO] Go application was not running or already stopped.")
    except Exception as e:
        print(f"[ERROR] An error occurred while running/monitoring the Go application: {e}")
        if go_process and go_process.poll() is None:
            go_process.kill()
        sys.exit(1)
    finally:
        exit_code = go_process.returncode if go_process and go_process.returncode is not None else 0
        print(f"\n[INFO] Go application exited with code {exit_code}.")
        print("[INFO] Script finished or interrupted.")
        print("[INFO] To stop the Docker containers, navigate to the 'infrastructure' directory and run: 'docker compose down' (or 'docker-compose down')")

if __name__ == "__main__":
    main()
