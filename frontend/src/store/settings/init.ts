import { addNotification } from '../notifications';
import { checkStatus } from '../status';
import {
  $settingsStore,
  changeSettingsFx,
  getSettingsFx,
  setSettings
} from './index';

$settingsStore
  .on(setSettings, ((_, settings) => settings));

getSettingsFx.doneData.watch(
  (result) => setSettings(result)
);

getSettingsFx.failData.watch(
  (error) => error.json().then((err) =>
    addNotification({
      type: 'error',
      prefix: 'serverError',
      message: err.error
    }))
);

changeSettingsFx.doneData.watch((result) => {
  setSettings(result);
  checkStatus();
});

changeSettingsFx.failData.watch(
  (error) => error.json().then((err) =>
    addNotification({
      type: 'error',
      prefix: 'serverError',
      message: err.error
    }))
);
