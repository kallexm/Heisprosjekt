//Elev_StateMachine.c

#include "Elev_StateMachine.h"
#include "Elev_Queue.h"
#include "elev.h"
#include "Elev_Timer.h"

#define ORDER_DIR_DOWN  0
#define ORDER_DIR_UP    1

#define BETWEEN_FLOORS -1


state_t state;
int lastFloor;
elev_Direction_type_t elevDirection;
int lastMove;
int hasBeenEmergency;



/*-----------------Private-----------------*/

static int abs(int number){
	if (number < 0){
		return number*-1;
	}
	else{
		return number;
	}
}




static void set_button_light(elev_button_type_t button, int floor, int value){
	if((floor == 3 && button == BUTTON_CALL_UP) || (floor == 0  && button == BUTTON_CALL_DOWN )){
	        return;
    	}

	else if(value == 0){
	 	if (button == BUTTON_CALL_UP && que_is_order_in_list(floor+1) == 1){
			return;
		} 
		else if(button == BUTTON_CALL_DOWN && que_is_order_in_list((floor+1)*-1) == 1){
	       		return;
		}
	}
       	elev_set_button_lamp(button,floor,value);
}




static void switch_all_floor_lamps(int floor, int value){
    set_button_light(BUTTON_CALL_UP,   floor, value);
    set_button_light(BUTTON_CALL_DOWN, floor, value);
    set_button_light(BUTTON_COMMAND,   floor, value);
}




static void turn_off_all_button_lamps(){
    	int i;
	for(i = 0; i <= 3; i++){
		switch_all_floor_lamps(i, 0);
	}
}




static void set_lastMove(){
	if(lastFloor == 4){
		lastMove = lastFloor;
	}
	else if( lastFloor == 1){
		lastMove = -1;
	}
	else{
		lastMove = lastFloor * elevDirection;
	}
}




static int order_interpreter(elev_button_type_t orderType, int orderFloor){
    if (orderType == BUTTON_CALL_UP){
		if(orderFloor==0){
            		return ORDER_DIR_DOWN;
		}
		else{
		    	return ORDER_DIR_UP;
		}
	}
	else if(orderType == BUTTON_CALL_DOWN){
		if(orderFloor == 3){
            		return ORDER_DIR_UP;
		}
		else{
		   	return ORDER_DIR_DOWN;
		}
	}
	else if(orderType == BUTTON_COMMAND){
		if (elevDirection == Down && orderFloor +1 >= lastFloor){
            		return ORDER_DIR_UP;
		}
		else if(elevDirection == Down && orderFloor + 1 < lastFloor){
			return ORDER_DIR_DOWN;
		}
		else if(elevDirection == Up && orderFloor + 1 > lastFloor){
		   	return ORDER_DIR_UP;
		}
		else if(elevDirection == Up && orderFloor + 1 <= lastFloor){
		    	return ORDER_DIR_DOWN;
		}
	}
	return -1;
}




static void add_order(int interpreter, int orderFloor){
    if(interpreter == ORDER_DIR_UP){
        que_add_order(orderFloor+1,lastMove);
    }
    else if(interpreter == ORDER_DIR_DOWN){
        que_add_order((orderFloor+1)*-1,lastMove);
    }
}




static void sm_state_initsializing(){
	elev_init();
	que_make_array();
	state = STATE_INITIALIZE;
	elev_set_motor_direction(DIRN_DOWN);
	elevDirection = Down;
}




static void sm_state_dorment(){
	state = STATE_DORMENT;
	elev_set_door_open_lamp(0);
	elev_set_motor_direction(DIRN_STOP);
}




static void sm_state_stop(){
	state = STATE_STOP;
	elev_set_motor_direction(DIRN_STOP);
	elev_set_door_open_lamp(1);
	time_start_timer(3);
}




static void sm_state_up(){
	state = STATE_UP;
	elev_set_door_open_lamp(0);
	elev_set_motor_direction(DIRN_UP);
	elevDirection = Up;
}




static void sm_state_down(){
	state = STATE_DOWN;
	elev_set_door_open_lamp(0);
	elev_set_motor_direction(DIRN_DOWN);
	elevDirection = Down;
}




static void sm_state_emergencystop(int floor){
    	state = STATE_EMERGENCY;
    	hasBeenEmergency = 1;

	elev_set_motor_direction(DIRN_STOP);
	time_reset_timer();
	que_del_all_order();
	turn_off_all_button_lamps();
	elev_set_stop_lamp(1);
	if(floor != BETWEEN_FLOORS){
		elev_set_door_open_lamp(1);
	}
}




static void set_elev_movestate(){
	int tempTopOrder = que_get_top_order();
	if(tempTopOrder == 0){
		sm_state_dorment();
	}
	else if (abs(tempTopOrder) < lastFloor){
		sm_state_down();
	}
	else if(abs(tempTopOrder) > lastFloor){
		sm_state_up();
	}
	else if(hasBeenEmergency == 1){
		if(tempTopOrder > 0){
			sm_state_up();
		}
		else if(tempTopOrder < 0){
			sm_state_down();
		}
		hasBeenEmergency = 0;
	}
}



/*-----------------Public-----------------*/

void sm_initiate(){
	sm_state_initsializing();
}




void sm_new_order(elev_button_type_t orderType, int orderFloor){
	if (state == STATE_INITIALIZE || state == STATE_EMERGENCY){
		return;
	}
	add_order(order_interpreter(orderType, orderFloor), orderFloor);
	set_button_light(orderType, orderFloor, 1);

	if (state != STATE_STOP){
		set_elev_movestate();
	}
}




void sm_emergency_stop(int emergencyButtonState, int floorSensor){
	if(state == STATE_INITIALIZE){
		return;
	}
	static int lastEmergencyButtonState = 0;
	if (emergencyButtonState == 1 && lastEmergencyButtonState != 1){
		sm_state_emergencystop(floorSensor);
	}
	else if (emergencyButtonState == 0 && lastEmergencyButtonState != 0){
		elev_set_stop_lamp(0);
		sm_state_dorment();
	}
	lastEmergencyButtonState = emergencyButtonState;
}



void sm_floor_sensor(int floor){
	if (floor != BETWEEN_FLOORS && state == STATE_INITIALIZE){
		lastFloor = floor +1;
		set_lastMove();
		sm_state_dorment();
	}
	else if(floor != BETWEEN_FLOORS){
		elev_set_floor_indicator(floor);
		lastFloor = floor +1;
		set_lastMove();
		if(floor + 1 == abs(que_get_top_order())){
			que_del_top_order();
            		switch_all_floor_lamps(floor, 0);
			sm_state_stop();
		}
	}
}




void sm_time_done(){
	time_reset_timer();
	set_elev_movestate();
}




