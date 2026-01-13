# Pier 🚢

Deploy anywhere. Command with AI. Own everything.

## Description

A next-generation deployment platform that combines the simplicity of Fly.io with the power of natural language AI. Deploy and manage infrastructure using plain English—no vendor lock-in, no complex commands to memorize.

## Why Pier?

Traditional cloud platforms make deployment easy but lock you into their pricing. Self-hosting gives you control but requires memorizing countless commands. Pier gives you both, plus AI-powered natural language control.

This Markdown version focuses exclusively on your **Phase 1** goals and the **MVP deadline** for March 31, 2026.

---

## 📍 Roadmap: Phase 1 (Jan – Mar 2026)

**Current Status:** ⚠️ Early Planning Stage

**Primary Goal:** Build a working CLI that deploys applications using natural language.

---

### 📅 Month 1: January 2026

*Focus: Core Infrastructure & AI Integration*

* **CLI Framework:** Build the base using **Go + Cobra**.
* **LLM Integration:** Connect **AWS Bedrock** for natural language processing.
* **Intent Recognition:** Map user prompts to specific actions (deploy, status, logs).
* **Provider Integration:** Build the **Hetzner Cloud API** client.
* **Initial Deployment:** Successful automated deployment to a single VPS.

### 📅 Month 2: February 2026

*Focus: Service Management & Proxying*

* **Containerization:** Full **Docker** container deployment support.
* **Networking:** Integrate **Caddy** for reverse proxying and automated **SSL certificates**.
* **Managed Services:** Natural language provisioning for **Postgres** and **Redis**.
* **Conversational Monitoring:** Implement commands to query system health via chat.
* **Contextual Memory:** Implement history tracking for "follow-up" commands.

### 📅 Month 3: March 2026

*Focus: Reliability & Polish*

* **Zero-Downtime:** Implement blue/green or rolling deployment strategies.
* **Safety Net:** Enable **natural language rollbacks** (e.g., "undo my last deploy").
* **Custom Domains:** Support for attaching user-owned domains.
* **AI Troubleshooting:** Use LLMs to diagnose and explain deployment failures.
* **Public Launch Prep:** Finalize documentation and record demo videos.

---

### 🚀 MVP Deliverables (March 31, 2026)

Upon completion of Phase 1, the following must be functional:

| Feature | Example Command |
| --- | --- |
| **App Deployment** | `pier deploy my-app` |
| **Database Provisioning** | `pier add postgres` |
| **Status Monitoring** | `pier show me status` |
| **Instant Rollback** | `pier undo that` |
| **Automated SSL** | *Handled automatically via Caddy* |

---

### 🎯 Success Criteria

* **Speed:** Deploy a real-world side project in **<2 minutes** using only natural language.
* **Accuracy:** AI correctly interprets **80%+** of commands without needing clarification.
* **Cost Efficiency:** Operating cost of **€4/mo** (hosting) + **~$5/mo** (Bedrock API).

---

Would you like me to draft a **GitHub README introductory section** based on these Phase 1 goals?

## License

See the [LICENSE](LICENSE) file for details.