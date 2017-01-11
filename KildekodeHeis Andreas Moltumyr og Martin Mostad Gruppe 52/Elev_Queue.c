//Elev_Queue.c

#define NUMBER_OF_FLOORS 4

int queueArrayLength =(NUMBER_OF_FLOORS-1)*2;

static int  queueArray[(NUMBER_OF_FLOORS-1)*2];
static int* queueTopPonter= queueArray;

static int  sortingArray[(NUMBER_OF_FLOORS-1)*2];
static int* sortingTopPointer = sortingArray;


/*-----------------Private-----------------*/

static void move_pointer(int** pointer, int* array){

	if (*pointer == &array[queueArrayLength-1]){
		*pointer = array;
	}
	else{
		(*pointer)++;
	}
}


static int get_index(int newOrder, int lastMove){
	int index;
	while(*sortingTopPointer != lastMove){
		move_pointer(&sortingTopPointer, sortingArray);
	}
	move_pointer(&sortingTopPointer, sortingArray);

	index = 0;
	while(*sortingTopPointer != newOrder){
		index ++;
		move_pointer(&sortingTopPointer, sortingArray);
	}
	return index;
}



/*-----------------Public-----------------*/

void que_make_array(){					// Fyller ut sortingArray etter NUMBER_OF_FLOORS
	int i;						// NUMBER_OF_FLOORS = 4 gir sortingArray = {-3, -2, -1, 2, 3, 4}
	for (i=0;i<queueArrayLength;i++){		// +/- angir retningen på orderen, nummer angir etasje som skal
		queueArray[i]=0;			// stoppes ved.
	}						// sortingArray benyttes i que_add_order() for å finne ut hvor
	for (i=0;i<(queueArrayLength);i++){		// en ny ordre skal legges i queueArray
		if (i<NUMBER_OF_FLOORS-1){
		sortingArray[i]=(i-NUMBER_OF_FLOORS+1);
		}
		else{
		sortingArray[i]=(i-NUMBER_OF_FLOORS+3);
		}
	}
}




int que_is_order_in_list(int newOrder){
	int i;
	for (i=0;i<queueArrayLength; i++){
        if (newOrder == queueArray[i]){
        return 1;
        }
    }
	return 0;
}




int que_get_top_order(){
	return *queueTopPonter;
}




void que_del_top_order(){
	*queueTopPonter = 0;
	move_pointer(&queueTopPonter, queueArray);
}




void que_del_all_order(){
	int i;
	for(i=0; i<queueArrayLength; i++){
		queueArray[i] = 0;
	}
	queueTopPonter = queueArray;
}




void que_add_order(int newOrder, int lastMove){
	int tempOrder;
	int* tempPointer;
	if(que_is_order_in_list(newOrder) != 0){
		return;
	}
	else if( newOrder == -4 || newOrder == 1){
		newOrder = newOrder*-1;	
	}
	tempPointer = queueTopPonter;
	do{
		if (*tempPointer == 0){
			*tempPointer = newOrder;
			return;
		}
		else if(get_index(newOrder, lastMove) < get_index(*tempPointer, lastMove)){
			tempOrder = *tempPointer;
			*tempPointer = newOrder;
			newOrder = tempOrder;
		}
		move_pointer(&tempPointer, queueArray);
	}while(queueTopPonter != tempPointer);
}




