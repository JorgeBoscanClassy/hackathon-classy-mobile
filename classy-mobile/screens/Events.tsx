import { StyleSheet, SafeAreaView, FlatList } from 'react-native';
import EditScreenInfo from '../components/EditScreenInfo';
import { Text, View } from '../components/Themed';
import { ListItem } from '../components/ListItems'
import { TouchableOpacity } from 'react-native-gesture-handler';
import {useState} from 'react'


export default function AttendessScreen({ navigation }) {

    const [ isFetching, setIsFetching ] = useState(false);

    const DATA = [
        {
          id: 1,
          title: 'Tunnels to Towers',
          label : '567 Attendees'
        },
        {
          id: 2,
          title: 'Los Angeles 5K',
          label : '145 Attendees'
        },
        {
          id: 3,
          title: 'Giving Tuesday 2022',
          label : '564 Attendees'
        }
      ];

      const renderItem = ({ item  }) => (
        <View key={item.id} style={{paddingLeft:15}}>
            <ListItem data={item} onPress={() => { navigation.navigate('Checkin') }} />
        </View>
      );


      const onRefresh = () => {

      }

      
    
  return (
    <View style={styles.container}>
      <View onPress={() => { navigation.navigate('Checkin') }}><Text style={styles.title}>Events</Text></View>
      <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
      <SafeAreaView style={styles.container}>
      <FlatList
        data={DATA}
        renderItem={renderItem}
        onRefresh={onRefresh}
        refreshing={isFetching}
        keyExtractor={item => item.id}
      />
    </SafeAreaView>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
  },
  title: {
    fontSize: 40,
    paddingLeft:15,
    paddingTop:20,
    paddingBottom:15,
    fontWeight: 'bold',
  },
});
