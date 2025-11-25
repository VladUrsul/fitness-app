from locust import HttpUser, task, between
import random
import datetime

class FitnessAppUser(HttpUser):
    wait_time = between(1, 1)  # simulate a user with 1s wait between tasks

    def on_start(self):
        """Run when a simulated user starts"""
        self.user_id = None
        self.workout_id = None
        self.session_id = None

    @task(2)
    def create_user(self):
        now = datetime.datetime.utcnow().isoformat() + "Z"
        payload = {
            "id": 0,
            "name": f"User_{random.randint(1, 1000000)}",
            "email": f"user_{random.randint(1, 1000000)}@example.com",
            "created_at": now,
            "updated_at": now
        }
        response = self.client.post("/api/v1/users", json=payload, name="/api/v1/users")
        if response.status_code == 200:
            self.user_id = response.json().get("id")

    @task(3)
    def create_workout(self):
        if not self.user_id:
            return
        now = datetime.datetime.utcnow().isoformat() + "Z"
        payload = {
            "id": 0,
            "type": f"Workout_{random.randint(1, 100)}",
            "user_id": self.user_id,
            "scheduled": now,
            "created_at": now,
            "updated_at": now
        }
        response = self.client.post("/api/v1/workouts", json=payload, name="/api/v1/workouts")
        if response.status_code == 200:
            self.workout_id = response.json().get("id")

    @task(4)
    def create_session(self):
        if not self.workout_id:
            return
        now = datetime.datetime.utcnow()
        payload = {
            "id": 0,
            "workout_id": self.workout_id,
            "started_at": now.isoformat() + "Z",
            "finished_at": (now + datetime.timedelta(minutes=random.randint(20, 90))).isoformat() + "Z",
            "created_at": now.isoformat() + "Z",
            "updated_at": now.isoformat() + "Z"
        }
        response = self.client.post("/api/v1/sessions", json=payload, name="/api/v1/sessions")
        if response.status_code == 200:
            self.session_id = response.json().get("id")

    @task(1)
    def get_user(self):
        if self.user_id:
            self.client.get(f"/api/v1/users/{self.user_id}", name="/api/v1/users/:id")

    @task(2)
    def get_workout(self):
        if self.workout_id:
            self.client.get(f"/api/v1/workouts/{self.workout_id}", name="/api/v1/workouts/:id")

    @task(3)
    def get_session(self):
        if self.session_id:
            self.client.get(f"/api/v1/sessions/{self.session_id}", name="/api/v1/sessions/:id")
