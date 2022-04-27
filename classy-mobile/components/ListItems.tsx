import React, { Component } from 'react';
import { StyleSheet, ScrollView, Dimensions } from 'react-native';
import { Text, View } from './Themed';
//import { Constants } from 'expo';

const { width } = Dimensions.get('window');

export class ListItems extends Component {
  
  componentDidMount() {
      
	}
	
  render() {
    return (
      <View 
        style={styles.container}
        //pagingEnabled={true}
        decelerationRate={0}
        snapToInterval={width - 60}
        snapToAlignment={"center"}
        contentInset={{
          top: 0,
          left: 0,
          bottom: 0,
          right: 0,
        }}>
            {this.props.data && this.props.data.map((card : any, index : number) => {

                return <View key={index} style={{flexDirection:'column'}}>
                    <View style={styles.view} >
                        <View style={{ marginRight: 'auto'}}>
                        <Text style={styles.title}>{card.title}</Text>
                        <Text style={styles.label}>{card.label}</Text>
                        </View>
                        <Text style={styles.right}>{card.right}</Text>
                        </View>
                        <View style={styles.separator} lightColor="#eee" darkColor="rgba(255,255,255,0.1)" />
                    </View>

            })}
      </View>
    );
  }
}

const styles = StyleSheet.create({
  container: {},
  view: {
    flexDirection:'row',
    margin: 0,
    width:width - 40,
    borderRadius: 10,
    paddingVertical:10,
    paddingHorizontal:0,
    alignContent:'center',
    justifyContent:'center'
  },
  label: {

    fontSize:14

  },
  title: {

    fontSize:20,
    marginBottom:3

  },
  right: {

    fontSize:30


  },
  separator: {
    height: 0,
    marginTop:5, 
    marginBottom:5,
    width: '100%',
  },
});

