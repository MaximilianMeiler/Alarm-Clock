import React from 'react';
import { View, Text, TextInput } from 'react-native';
import { useDarkMode } from '../contexts/DarkModeContext';
import { getDynamicStyles } from '../styles/AlarmStyles';

function AlarmNameInput({ alarmName, setAlarmName }) {
  const { isDarkMode } = useDarkMode();
  const dynamicStyles = getDynamicStyles(isDarkMode); 

  return (
    <View style={dynamicStyles.container}> 
      <Text style={dynamicStyles.label}>Label</Text>
      <TextInput
        style={dynamicStyles.textInput} // Apply the dynamic styles here too 
        onChangeText={setAlarmName}
        value={alarmName}
        placeholder="Alarm"
      />
    </View>
  );
}

export default AlarmNameInput;
