Here are messaging services that could be added to msgr:

  Chat & Team Communication
  ┌─────────────────┬──────────────────────┬───────────────────────────────┐
  │     Service     │       API Type       │           Use Case            │
  ├─────────────────┼──────────────────────┼───────────────────────────────┤
  │ Discord         │ Webhooks / Bot API   │ Gaming communities, dev teams │
  ├─────────────────┼──────────────────────┼───────────────────────────────┤
  │ Microsoft Teams │ Webhooks / Graph API │ Enterprise                    │
  ├─────────────────┼──────────────────────┼───────────────────────────────┤
  │ Mattermost      │ Webhooks             │ Self-hosted Slack alternative │
  ├─────────────────┼──────────────────────┼───────────────────────────────┤
  │ Rocket.Chat     │ REST API             │ Self-hosted team chat         │
  ├─────────────────┼──────────────────────┼───────────────────────────────┤
  │ Zulip           │ REST API             │ Open source, threaded chat    │
  ├─────────────────┼──────────────────────┼───────────────────────────────┤
  │ Google Chat     │ Webhooks             │ Google Workspace teams        │
  ├─────────────────┼──────────────────────┼───────────────────────────────┤
  │ Cisco Webex     │ Webhooks / API       │ Enterprise                    │
  ├─────────────────┼──────────────────────┼───────────────────────────────┤
  │ Flock           │ Webhooks             │ Team messaging                │
  ├─────────────────┼──────────────────────┼───────────────────────────────┤
  │ Twist           │ REST API             │ Async team communication      │
  └─────────────────┴──────────────────────┴───────────────────────────────┘
  Push Notifications
  ┌──────────────────────────┬──────────┬─────────────────────────────┐
  │         Service          │ API Type │          Use Case           │
  ├──────────────────────────┼──────────┼─────────────────────────────┤
  │ Pushover                 │ REST API │ Mobile push notifications   │
  ├──────────────────────────┼──────────┼─────────────────────────────┤
  │ Pushbullet               │ REST API │ Cross-device notifications  │
  ├──────────────────────────┼──────────┼─────────────────────────────┤
  │ Ntfy                     │ REST API │ Self-hosted push, no signup │
  ├──────────────────────────┼──────────┼─────────────────────────────┤
  │ Gotify                   │ REST API │ Self-hosted push server     │
  ├──────────────────────────┼──────────┼─────────────────────────────┤
  │ Simplepush               │ REST API │ Encrypted push              │
  ├──────────────────────────┼──────────┼─────────────────────────────┤
  │ Pusher                   │ REST API │ Real-time notifications     │
  ├──────────────────────────┼──────────┼─────────────────────────────┤
  │ OneSignal                │ REST API │ Mass push notifications     │
  ├──────────────────────────┼──────────┼─────────────────────────────┤
  │ Firebase Cloud Messaging │ REST API │ Android/iOS/Web push        │
  └──────────────────────────┴──────────┴─────────────────────────────┘
  Email Services
  ┌────────────────┬───────────────┬──────────────────────┐
  │    Service     │   API Type    │       Use Case       │
  ├────────────────┼───────────────┼──────────────────────┤
  │ SendGrid       │ REST API      │ Transactional email  │
  ├────────────────┼───────────────┼──────────────────────┤
  │ AWS SES        │ AWS SDK       │ Scalable email       │
  ├────────────────┼───────────────┼──────────────────────┤
  │ Postmark       │ REST API      │ Transactional email  │
  ├────────────────┼───────────────┼──────────────────────┤
  │ Mailjet        │ REST API      │ Email delivery       │
  ├────────────────┼───────────────┼──────────────────────┤
  │ SparkPost      │ REST API      │ Email infrastructure │
  ├────────────────┼───────────────┼──────────────────────┤
  │ Resend         │ REST API      │ Modern email API     │
  ├────────────────┼───────────────┼──────────────────────┤
  │ SMTP (generic) │ SMTP protocol │ Any mail server      │
  └────────────────┴───────────────┴──────────────────────┘
  SMS & Voice
  ┌────────────────┬──────────┬────────────────┐
  │    Service     │ API Type │    Use Case    │
  ├────────────────┼──────────┼────────────────┤
  │ Vonage (Nexmo) │ REST API │ SMS/Voice      │
  ├────────────────┼──────────┼────────────────┤
  │ Plivo          │ REST API │ SMS/Voice      │
  ├────────────────┼──────────┼────────────────┤
  │ MessageBird    │ REST API │ SMS/Voice/Chat │
  ├────────────────┼──────────┼────────────────┤
  │ Sinch          │ REST API │ SMS/Voice      │
  ├────────────────┼──────────┼────────────────┤
  │ AWS SNS        │ AWS SDK  │ SMS at scale   │
  ├────────────────┼──────────┼────────────────┤
  │ Clickatell     │ REST API │ Enterprise SMS │
  ├────────────────┼──────────┼────────────────┤
  │ Textbelt       │ REST API │ Simple SMS     │
  └────────────────┴──────────┴────────────────┘
  Webhooks & Automation
  ┌───────────────────┬───────────┬────────────────────────┐
  │      Service      │ API Type  │        Use Case        │
  ├───────────────────┼───────────┼────────────────────────┤
  │ Generic Webhook   │ HTTP POST │ Any webhook endpoint   │
  ├───────────────────┼───────────┼────────────────────────┤
  │ IFTTT             │ Webhooks  │ Automation triggers    │
  ├───────────────────┼───────────┼────────────────────────┤
  │ Zapier            │ Webhooks  │ Workflow automation    │
  ├───────────────────┼───────────┼────────────────────────┤
  │ Make (Integromat) │ Webhooks  │ Automation             │
  ├───────────────────┼───────────┼────────────────────────┤
  │ n8n               │ Webhooks  │ Self-hosted automation │
  └───────────────────┴───────────┴────────────────────────┘
  Developer & DevOps
  ┌────────────────────┬─────────────┬────────────────────────┐
  │      Service       │  API Type   │        Use Case        │
  ├────────────────────┼─────────────┼────────────────────────┤
  │ PagerDuty          │ REST API    │ Incident alerts        │
  ├────────────────────┼─────────────┼────────────────────────┤
  │ Opsgenie           │ REST API    │ Alert management       │
  ├────────────────────┼─────────────┼────────────────────────┤
  │ VictorOps (Splunk) │ REST API    │ Incident response      │
  ├────────────────────┼─────────────┼────────────────────────┤
  │ Datadog Events     │ REST API    │ Monitoring events      │
  ├────────────────────┼─────────────┼────────────────────────┤
  │ New Relic          │ REST API    │ APM notifications      │
  ├────────────────────┼─────────────┼────────────────────────┤
  │ Sentry             │ REST API    │ Error tracking         │
  ├────────────────────┼─────────────┼────────────────────────┤
  │ GitHub Issues      │ REST API    │ Create issues from CLI │
  ├────────────────────┼─────────────┼────────────────────────┤
  │ GitLab Issues      │ REST API    │ Create issues from CLI │
  ├────────────────────┼─────────────┼────────────────────────┤
  │ Jira               │ REST API    │ Create tickets         │
  ├────────────────────┼─────────────┼────────────────────────┤
  │ Linear             │ GraphQL API │ Create issues          │
  └────────────────────┴─────────────┴────────────────────────┘
  Social Media
  ┌────────────────┬─────────────┬──────────────────────┐
  │    Service     │  API Type   │       Use Case       │
  ├────────────────┼─────────────┼──────────────────────┤
  │ Twitter/X      │ REST API    │ Post tweets          │
  ├────────────────┼─────────────┼──────────────────────┤
  │ Mastodon       │ REST API    │ Fediverse posts      │
  ├────────────────┼─────────────┼──────────────────────┤
  │ Bluesky        │ AT Protocol │ Post to Bluesky      │
  ├────────────────┼─────────────┼──────────────────────┤
  │ LinkedIn       │ REST API    │ Professional updates │
  ├────────────────┼─────────────┼──────────────────────┤
  │ Facebook Pages │ Graph API   │ Page posts           │
  └────────────────┴─────────────┴──────────────────────┘
  Chat Apps (Consumer)
  ┌───────────────────┬──────────────────────┬────────────────────┐
  │      Service      │       API Type       │      Use Case      │
  ├───────────────────┼──────────────────────┼────────────────────┤
  │ WhatsApp Business │ Cloud API            │ Business messaging │
  ├───────────────────┼──────────────────────┼────────────────────┤
  │ Signal            │ Signal CLI           │ Secure messaging   │
  ├───────────────────┼──────────────────────┼────────────────────┤
  │ Viber             │ REST API             │ Viber bots         │
  ├───────────────────┼──────────────────────┼────────────────────┤
  │ LINE              │ Messaging API        │ Popular in Asia    │
  ├───────────────────┼──────────────────────┼────────────────────┤
  │ WeChat            │ Official Account API │ China market       │
  ├───────────────────┼──────────────────────┼────────────────────┤
  │ KakaoTalk         │ REST API             │ Korea market       │
  └───────────────────┴──────────────────────┴────────────────────┘
  Matrix Protocol
  ┌───────────────────┬────────────┬────────────────────┐
  │      Service      │  API Type  │      Use Case      │
  ├───────────────────┼────────────┼────────────────────┤
  │ Matrix (Element)  │ Matrix API │ Decentralized chat │
  ├───────────────────┼────────────┼────────────────────┤
  │ Any Matrix server │ Matrix API │ Self-hosted        │
  └───────────────────┴────────────┴────────────────────┘
  IRC
  ┌───────────────┬──────────────┬─────────────────────────┐
  │    Service    │   API Type   │        Use Case         │
  ├───────────────┼──────────────┼─────────────────────────┤
  │ IRC (generic) │ IRC protocol │ Legacy systems          │
  ├───────────────┼──────────────┼─────────────────────────┤
  │ Libera.Chat   │ IRC protocol │ Open source communities │
  └───────────────┴──────────────┴─────────────────────────┘
  IoT & Hardware
  ┌────────────────┬───────────────┬──────────────────────────┐
  │    Service     │   API Type    │         Use Case         │
  ├────────────────┼───────────────┼──────────────────────────┤
  │ Home Assistant │ REST API      │ Smart home notifications │
  ├────────────────┼───────────────┼──────────────────────────┤
  │ MQTT           │ MQTT protocol │ IoT messaging            │
  └────────────────┴───────────────┴──────────────────────────┘
  Quick Wins (Easiest to Implement)

  1. Discord - Simple webhooks, very popular
  2. Pushover - Dead simple REST API
  3. Ntfy - No auth required, self-hostable
  4. Generic Webhook - POST JSON anywhere
  5. SendGrid - Popular email alternative to Mailgun
  6. Matrix - Decentralized, growing adoption
