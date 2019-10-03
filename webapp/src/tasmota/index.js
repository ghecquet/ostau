import PahoMQTT from 'paho-mqtt'
import MQTTPattern from 'mqtt-pattern'

import { temperatureSetting } from '../actions/traits';

const reduxMqttMiddleware = config => ({ dispatch }) => {
    const client = new PahoMQTT.Client(config.host, config.port, config.clientId || "");

    const statusPattern = 'stat/+id/+cmd'
    const sensorPattern = 'tele/+id/SENSOR'

    client.connect({onSuccess: () => {
        client.subscribe(MQTTPattern.clean(statusPattern));
        client.subscribe(MQTTPattern.clean(sensorPattern));
    }})

    client.onMessageArrived = ((message) => {
        const statusParams = MQTTPattern.exec(statusPattern, message.topic)
        if (statusParams) {

        }

        const sensorParams = MQTTPattern.exec(sensorPattern, message.topic)
        if (sensorParams) {
            const value = JSON.parse(message.payloadString);

            dispatch(temperatureSetting({
                id: sensorParams.id,
                type: "DEVICE",
                temperature: value.SI7021.Temperature,
                humidity: value.SI7021.Humidity
            }))
        }
    })

//   client.on('message', ((topic, message) => {
//     const msgObj = JSON.parse(message);
//     if (topic === 'unregister') {
//       client.unsubscribe(msgObj.deviceId);
//       dispatch(removeDevice(msgObj));
//     } else if (topic === 'registrations') {
//       client.subscribe(msgObj.deviceId);
//       dispatch(addDevice(msgObj));
//     } else {
//       console.log('message:', msgObj);
//       dispatch(onData(msgObj));
//     }
//   }));

    return next => (action) => {
        console.log('nextaction', action);
        next(action);
    };
};

export default reduxMqttMiddleware;