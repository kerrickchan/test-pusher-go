import * as PusherPushNotifications from "@pusher/push-notifications-web";

const beamsClient = new PusherPushNotifications.Client({
  instanceId: import.meta.env.VITE_PUSHER_BEAMS_INSTANCE_ID,
});

beamsClient
  .start()
  .then(() => beamsClient.getDeviceId())
  .then((deviceId) =>
    console.log("Successfully registered with Beams. Device ID:", deviceId)
  )
  .then(() => beamsClient.addDeviceInterest("debug-hello"))
  .then(() => beamsClient.getDeviceInterests())
  .then((interests) => console.log("Current interests:", interests))
  .catch(console.error);
