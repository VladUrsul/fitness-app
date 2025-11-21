from locust import HttpUser, task, between
import random
import datetime
import time

def unique_email():
    return f"user_{int(time.time() * 1000)}_{random.randint(10000, 99999)}@example.com"

def unique_name():
    return f"User_{int(time.time() * 1000)}_{random.randint(10000, 99999)}"

def random_time():
    return datetime.datetime.utcnow().isoformat() + "Z"

class FitnessMonolithUser(HttpUser):
    wait_time = between(1, 1)

    @task
    def workflow(self):
        # 1. Create user
        user_payload = {
            "name": unique_name(),
            "email": unique_email()
        }
        user_resp = self.client.post(
            "/api/v1/users",
            json=user_payload,
            name="Monolith - Create User"
        )
        if user_resp.status_code != 200:
            return

        user_id = user_resp.json().get("id")

        # 2. Create workout
        workout_payload = {
            "user_id": user_id,
            "type": random.choice(["Yoga", "Cardio", "Strength"]),
            "scheduled": random_time()
        }
        workout_resp = self.client.post(
            "/api/v1/workouts",
            json=workout_payload,
            name="Monolith - Create Workout"
        )
        if workout_resp.status_code != 200:
            return

        workout_id = workout_resp.json().get("id")

        # 3. Create session
        session_payload = {
            "workout_id": workout_id,
            "started_at": random_time(),
            "finished_at": random_time()
        }
        self.client.post(
            "/api/v1/sessions",
            json=session_payload,
            name="Monolith - Create Session"
        )
