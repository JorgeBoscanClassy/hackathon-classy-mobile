import React, { Component } from 'react';
import { View, StyleSheet, ScrollView, Dimensions, Text } from 'react-native';
//import { Constants } from 'expo';

const { width } = Dimensions.get('window');

export class Highlights extends Component {
  
  componentDidMount() {
		setTimeout(() => {this.scrollView.scrollTo({x: -30}) }, 1) // scroll view position fix
	}
	
  render() {
    return (
      <ScrollView 
        ref={(scrollView) => { this.scrollView = scrollView; }}
        style={styles.container}
        //pagingEnabled={true}
        showsHorizontalScrollIndicator={false}
        horizontal= {true}
        decelerationRate={0}
        snapToInterval={width - 60}
        snapToAlignment={"center"}
        contentInset={{
          top: 0,
          left: 0,
          bottom: 0,
          right: 30,
        }}>
            {this.props.data && this.props.data.map((card : any, index : number) => {

                return <View key={index} style={styles.view} >
                        <Text style={styles.title}>{card.title}</Text>
                        <Text style={styles.label}>{card.label}</Text>
                        </View>

            })}
      </ScrollView>
    );
  }
}

const styles = StyleSheet.create({
  container: {},
  view: {
    marginTop: 20,
    backgroundColor: '#EB7251',
    width: 250,
    margin: 10,
    marginLeft:0,
    height: 100,
    borderRadius: 10,
    paddingVertical:20,
    paddingHorizontal:20
  },
  label: {

    color:'#fff',
    fontSize:15

  },
  title: {

    color:'#fff',
    fontWeight:'bold',
    fontSize:30


  }
});

