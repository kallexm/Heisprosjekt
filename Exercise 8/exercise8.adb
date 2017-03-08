with Ada.Text_IO, Ada.Integer_Text_IO, Ada.Numerics.Float_Random;
use  Ada.Text_IO, Ada.Integer_Text_IO, Ada.Numerics.Float_Random;

procedure exercise8 is

    Count_Failed    : exception;    -- Exception to be raised when counting fails
    Gen             : Generator;    -- Random number generator

    protected type Transaction_Manager (N : Positive) is
        entry Finished;
        entry Wait_Until_Aborted;
        procedure Signal_Abort;
    private
        Finished_Gate_Open  : Boolean := False;
        Aborted             : Boolean := False;
    end Transaction_Manager;
    protected body Transaction_Manager is
        entry Finished when Finished_Gate_Open or Finished'Count = N is
        begin
            ------------------------------------------
            -- PART 3: Complete the exit protocol here
			Finished_Gate_Open := True;
			if Finished'Count = 0 then
				Finished_Gate_Open := false;
			end if;
            ------------------------------------------
        end Finished;

        entry Wait_Until_Aborted when Aborted is
        begin
            if Wait_Until_Aborted'Count = 0 then
                Aborted := False;
            end if;
        end Wait_Until_Aborted;

        procedure Signal_Abort is
        begin
            Aborted := True;
        end Signal_Abort;

    end Transaction_Manager;



    
    function Unreliable_Slow_Add (x : Integer) return Integer is
    Error_Rate : Constant := 0.15;  -- (between 0 and 1)
    d : float;
	begin
        -------------------------------------------
        -- PART 1: Create the transaction work here
		if Random(Gen) >= Error_Rate then
			d := 3.50 + 1.0*Random(Gen);
			delay Duration(d);
			return x+10;
		else
			d := 0.50 + 0.50*Random(Gen);
            delay Duration(d);
			raise Count_Failed;
		end if;
        -------------------------------------------
    end Unreliable_Slow_Add;




    task type Transaction_Worker (Initial : Integer; Manager : access Transaction_Manager);
    task body Transaction_Worker is
        Num         : Integer   := Initial;
        Round_Num   : Integer   := 0;
    begin
        Put_Line ("Worker" & Integer'Image(Initial) & " started");

        loop
            Put_Line ("Worker" & Integer'Image(Initial) & " started round" & Integer'Image(Round_Num));
            Round_Num := Round_Num + 1;

            ---------------------------------------
            -- PART 2: Do the transaction work here
            begin	
                select
                    Manager.Wait_Until_Aborted;
                    num := num + 5;
                    Put_Line ("Forward Error Recovery!");
                then abort
                    begin
                        Num := Unreliable_Slow_Add(Num);
                    exception
                        when Count_Failed =>
                            Manager.Signal_Abort;
                        when others =>
                            Put_Line ("Some Error Accured!");
                    end;
                    Manager.Finished;
                end select;
                
            end;

            ---------------------------------------
            
            
            Put_Line ("  Worker" & Integer'Image(Initial) & " comitting" & Integer'Image(Num));

            delay 0.5;

        end loop;
    end Transaction_Worker;

    Manager : aliased Transaction_Manager (3);

    Worker_1 : Transaction_Worker (0, Manager'Access);
    Worker_2 : Transaction_Worker (1, Manager'Access);
    Worker_3 : Transaction_Worker (2, Manager'Access);

begin
    Reset(Gen); -- Seed the random number generator
end exercise8;
